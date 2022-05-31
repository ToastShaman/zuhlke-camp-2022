import org.junit.jupiter.api.Test
import strikt.api.expectThat
import strikt.assertions.containsExactly
import strikt.assertions.first
import strikt.assertions.map
import java.util.*

class SocialNetworkKataTest {

    @Test
    fun `Posting - Alice can publish messages to a personal timeline`() {
        val network = TheSocialNetwork()

        val alice = network.createUser(Username("Alice"))
            .post(Message("I am cooking breakfast!"))

        expectThat(alice.timeline.allMessages())
            .map(TimelineMessage::message)
            .containsExactly(Message("I am cooking breakfast!"))
    }

    @Test
    fun `Reading - Bob can view Alice's timeline`() {
        val network = TheSocialNetwork()

        val alice = network.createUser(Username("Alice"))
            .post(Message("I am cooking breakfast!"))

        val bob = network.createUser(Username("Bob"))

        val timeline = bob.viewTimelineFrom(alice)

        expectThat(timeline.allMessages())
            .map(TimelineMessage::message)
            .containsExactly(Message("I am cooking breakfast!"))
    }

    @Test
    fun `Following - Charlie can subscribe to Alice's and Bob's timelines, and view an aggregated list of all subscriptions`() {
        val network = TheSocialNetwork()

        val alice = network.createUser(Username("Alice"))
            .post(Message("I am cooking breakfast!"))

        val bob = network.createUser(Username("Bob"))
            .post(Message("I am cooking dinner!"))

        val charlie = network.createUser(Username("Charlie"))

        val timeline = charlie.subscribeTo(alice, bob)

        expectThat(timeline.allMessages())
            .containsExactly(
                TimelineMessage(alice, Message("I am cooking breakfast!")),
                TimelineMessage(bob, Message("I am cooking dinner!")),
            )
    }

    @Test
    fun `Mentions - Bob can link to Charlie in a message using @`() {
        val network = TheSocialNetwork()

        val charlie = network.createUser(Username("Charlie"))

        val bob = network.createUser(Username("Bob"))
            .post(Message("I am cooking dinner! Come over @Charlie"))

        expectThat(charlie.viewTimelineFrom(bob))
            .get { message }
            .first()
            .get { mentions }
            .containsExactly(charlie)
    }
}

class TheSocialNetwork {
    private val users = mutableMapOf<Username, User>()

    fun createUser(username: Username) = users
        .computeIfAbsent(username) {
            User(
                username = username,
                timeline = MutableTimeline(),
                network = this
            )
        }

    fun find(username: Username) = users[username]
}

@JvmInline
value class Username(val value: String)

@JvmInline
value class Message(val value: String) {
    init {
        require(value.isNotBlank())
    }

    fun maybeMentions() = "@([a-zA-Z\\d]+)"
        .toRegex()
        .find(value)
        ?.groupValues
        ?.get(1)
        ?.let(::Username)
}

data class TimelineMessage(
    val user: User,
    val message: Message,
    val mentions: List<User> = emptyList()
)

interface Timeline {
    fun allMessages(): List<TimelineMessage>
}

class MutableTimeline : Timeline {
    private val messages = LinkedList<TimelineMessage>()

    fun post(message: TimelineMessage) {
        messages.addFirst(message)
    }

    override fun allMessages() = messages.toList()
}

data class ImmutableTimeline(val message: List<TimelineMessage>) : Timeline {
    override fun allMessages() = message
}

data class User(
    val username: Username,
    val timeline: MutableTimeline,
    val network: TheSocialNetwork
) {
    fun post(message: Message) = apply {
        timeline.post(
            TimelineMessage(
                user = this,
                message = message,
                mentions = message.maybeMentions()?.let(network::find)?.let(::listOf).orEmpty()
            )
        )
    }

    fun viewTimelineFrom(other: User) = ImmutableTimeline(other.timeline.allMessages())

    fun subscribeTo(vararg users: User) = users
        .map(User::timeline)
        .flatMap(MutableTimeline::allMessages)
        .let(::ImmutableTimeline)
}
