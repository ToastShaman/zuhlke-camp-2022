# A Pairing Session Template

Source: https://tuple.app/pair-programming-guide/template

1. [ ] Agree on the high-level goal out loud.

1. [ ] Break the work into a handful of tasks and prioritise them.

1. [ ] Decide your driver/navigator swapping strategy.

1. [ ] Configure git to share credit.

1. [ ] Eliminate distractions.

1. [ ] Work.

1. [ ] Analyse the session with a mini retro.

## Agree on the high-level goal out loud

State out loud what you hope to accomplish at a high level.

You wouldn't think it'd be possible for two people to start pairing without agreement about where they're headed, but it's surprisingly easy.

## Break the work into a handful of tasks (and prioritise them)

It's worth trying to break your high-level goal into a handful of smaller steps.

This has a number of benefits:

* It makes the goal less intimidating.

* You'll spot dead ends and pitfalls more easily.

* You can sort your task list by priority.

* You're more likely to notice that accomplishing task C would make B easier, and reorder appropriately.

* You can decide on a task based on your current energy levels.

* It gives you a clear place to put new tasks you think of while working.

[Some folks][1] like to write each task on its own index card. The stack of them lives in front of the navigator. Each card can be a nice home for notes or ideas to bring up when there is a break in the action.

## Decide what will trigger a driver/navigator swap

Unless you already know what works best for you, I strongly recommend the [Pomodoro Technique][2]:

1. Code for 25 minutes.

1. Take a 5 minute break.

1. Switch drivers.

Other [pair programming styles][3] exist if you wish to try them.

## Configure git to share credit

If two of you work on some code, both your names should appear on the commit.

Here's a [handy guide][4] to configuring git appropriately.

Bonus: GitHub understands this natively and will give you both credit for the commit.

A few tools exist to make this even easier:

* [git pair](https://github.com/chrisk/git-pair)

* [git duet](https://github.com/git-duet/git-duet)

* [git-together](https://github.com/kejadlen/git-together)

## Eliminate distractions

Show respect for your pair and the work you're about to do.

* Don't bring your phone. Silence it if you do.

* Disable notifications on the machine you're using to pair.

* Close email/Slack/Twitter/IRC. Never keep something distracting on a second monitor.

## Work

Do the work!

Don't forget:

* *When navigating*: ask questions rather than making demands.
* *When driving*: dictate what you're doing and why.
* Err on the side of over-communication.
* Take lots of breaks.
* Swap roles frequently.
* Do the simplest thing that could possibly work (for now).
* Avoid these pairing [anti-patterns][5].

## Perform a mini retro

Spend a few minutes after your session reflecting on the experience.

First, discuss what went well.

Then, consider what would make the next session 1% better.

Possible areas for improvement:

* **Focus**: did distractions sneak in?

* **Communication**: were there long stretches of no talking?

* **Pacing**: did the session feel like a grind at any point?

* **Division of responsibility**: did you split the work up well?

* **Code quality**: was your end-product high-quality?

[1]: https://www.jamesshore.com/Agile-Book/pair_programming.html
[2]: https://en.wikipedia.org/wiki/Pomodoro_Technique
[3]: https://tuple.app/pair-programming-guide/styles
[4]: https://help.github.com/articles/creating-a-commit-with-multiple-authors/
[5]: https://tuple.app/pair-programming-guide/antipatterns