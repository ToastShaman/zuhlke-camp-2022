import org.junit.jupiter.api.Test
import strikt.api.expectThat
import strikt.assertions.containsExactly
import java.util.*

class SocialNetworkKataTest {

    @Test
    fun `Posting - Alice can publish messages to a personal timeline`() {
        val alice = User(Username("Alice"), MutableTimeline())
            .post(Message("I am cooking breakfast!"))

        expectThat(alice.timeline.allMessages())
            .containsExactly(Message("I am cooking breakfast!"))
    }

    @Test
    fun `Reading - Bob can view Alice timeline`() {
        val alice = User(Username("Alice"), MutableTimeline())
            .post(Message("I am cooking breakfast!"))

        val bob = User(Username("Bob"), MutableTimeline())

        val timeline = bob.viewTimelineFrom(alice)

        expectThat(timeline.allMessages())
            .containsExactly(Message("I am cooking breakfast!"))
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

interface Timeline {
    fun allMessages(): List<Message>
}

class MutableTimeline : Timeline {
    private val messages: LinkedList<Message> = LinkedList()

    fun post(message: Message) {
        messages.addFirst(message)
    }

    override fun allMessages() = messages.toList()
}

data class ImmutableTimeline(val message: List<Message>) : Timeline {
    override fun allMessages() = message
}

data class User(
    val username: Username,
    val timeline: MutableTimeline
) {
    fun post(message: Message) = apply {
        timeline.post(message)
    }

    fun viewTimelineFrom(other: User) = ImmutableTimeline(other.timeline.allMessages())
}
