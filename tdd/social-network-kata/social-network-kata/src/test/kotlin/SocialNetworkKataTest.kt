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
}

@JvmInline
value class Username(val value: String)

@JvmInline
value class Message(val value: String) {
    init {
        require(value.isNotBlank())
    }
}

class MutableTimeline {
    private val messages: LinkedList<Message> = LinkedList()

    fun post(message: Message) {
        messages.addFirst(message)
    }

    fun allMessages() = messages.toList()
}

data class User(
    val username: Username,
    val timeline: MutableTimeline
) {
    fun post(message: Message) = apply {
        timeline.post(message)
    }
}
