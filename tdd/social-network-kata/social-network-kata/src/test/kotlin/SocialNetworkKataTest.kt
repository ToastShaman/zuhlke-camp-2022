import org.junit.jupiter.api.Test
import strikt.api.expectThat
import strikt.assertions.containsExactly
import strikt.assertions.map
import java.util.*

class SocialNetworkKataTest {

    @Test
    fun `Posting - Alice can publish messages to a personal timeline`() {
        val alice = User(Username("Alice"), MutableTimeline())
            .post(Message("I am cooking breakfast!"))

        expectThat(alice.timeline.allMessages())
            .map(TimelineMessage::message)
            .containsExactly(Message("I am cooking breakfast!"))
    }

    @Test
    fun `Reading - Bob can view Alice's timeline`() {
        val alice = User(Username("Alice"), MutableTimeline())
            .post(Message("I am cooking breakfast!"))

        val bob = User(Username("Bob"), MutableTimeline())

        val timeline = bob.viewTimelineFrom(alice)

        expectThat(timeline.allMessages())
            .map(TimelineMessage::message)
            .containsExactly(Message("I am cooking breakfast!"))
    }

    @Test
    fun `Following - Charlie can subscribe to Alice's and Bob's timelines, and view an aggregated list of all subscriptions`() {
        val alice = User(Username("Alice"), MutableTimeline())
            .post(Message("I am cooking breakfast!"))

        val bob = User(Username("Bob"), MutableTimeline())
            .post(Message("I am cooking dinner!"))

        val charlie = User(Username("Charlie"), MutableTimeline())

        val timeline = charlie.subscribeTo(alice, bob)

        expectThat(timeline.allMessages())
            .containsExactly(
                TimelineMessage(alice, Message("I am cooking breakfast!")),
                TimelineMessage(bob, Message("I am cooking dinner!")),
            )
    }
}

@JvmInline
value class Username(val value: String)

@JvmInline
value class Message(val value: String) {
    init {
        require(value.isNotBlank())
    }
}

data class TimelineMessage(
    val user: User,
    val message: Message
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
    val timeline: MutableTimeline
) {
    fun post(message: Message) = apply {
        timeline.post(TimelineMessage(this, message))
    }

    fun viewTimelineFrom(other: User) = ImmutableTimeline(other.timeline.allMessages())

    fun subscribeTo(vararg users: User) = users
        .map(User::timeline)
        .flatMap(MutableTimeline::allMessages)
        .let(::ImmutableTimeline)
}
