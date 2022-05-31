import org.junit.jupiter.api.Test
import strikt.api.expectThat
import strikt.assertions.containsExactly
import strikt.assertions.first
import strikt.assertions.map
import java.net.URL
import java.util.*

/**
 * https://kata-log.rocks/social-network-kata
 */
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
            .get(ImmutableTimeline::message)
            .first()
            .get(TimelineMessage::mentions)
            .containsExactly(charlie)
    }

    @Test
    fun `Links - Alice can link to a clickable web resource in a message`() {
        val network = TheSocialNetwork()

        val alice = network.createUser(Username("Alice"))
            .post(Message("Look at this video", Link(URL("https://bit.ly/3M8JUPv"))))

        expectThat(alice.timeline)
            .get(MutableTimeline::allMessages)
            .first()
            .get(TimelineMessage::links)
            .containsExactly(Link(URL("https://bit.ly/3M8JUPv")))
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
value class Link(val url: URL)

data class Message(
    val value: String,
    val link: Link? = null
) {
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
    val mentions: List<User> = emptyList(),
    val links: List<Link> = emptyList(),
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
                mentions = message.maybeMentions()?.let(network::find)?.let(::listOf).orEmpty(),
                links = message.link?.let(::listOf).orEmpty()
            )
        )
    }

    fun viewTimelineFrom(other: User) = ImmutableTimeline(other.timeline.allMessages())

    fun subscribeTo(vararg users: User) = users
        .map(User::timeline)
        .flatMap(MutableTimeline::allMessages)
        .let(::ImmutableTimeline)
}
