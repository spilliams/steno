- 2020-11-20 I adopted a nav/command dictionary that I found on the internet.
    http://www.openstenoproject.org/stenodict/dictionaries/single_stroke_commands.html
    I have by now also read the section in Learn Plover! about dictionaries, so
    I feel well-equipped to start tabbing my way around.  
    The only real trouble is (on a Windows machine) I keep this repo
    checked-out to a directory in WSL2, which is not (easily) accessible from
    Windows, where Plover is installed. Whoops.  
    Ok, I realized (again) that this will not solve all my nav use cases. I
    still need a way to alt-tab that either lets me hold alt while hitting tab
    any number of times (and sending those inputs individually, so I can see
    results), or I need a brief for "hit alt-tab twice in a row" etc, up to
    like 4 or 5 (the max number of windows I'd want to be open in a space
    anyway). Gonna try setting some of my own nav briefs for this.  
    <experiments with Alt_L(Tab Tab Tab) ensue>  
    I think this particular behavior is not something I can achieve with steno.
    Perhaps it is time to back up and take another look at the QMK symbol layer
    bound to my bottom left pinky key. Another night.
- 2020-11-19 Today I updated my qmk_firmware repo to be updated with what Germ
    has on their side. I tried a PR, but that got really mucky with a changed
    file and a changed submodule, so I backed out, saved patches of my commits,
    and completely remade my fork. It turns out the changed file also hit my
    new fork (wtf?), but the patches applied just fine.  
    Last night I read about dictionaries, and how "Learn Plover" teaches the
    shift/ctrl/alt/super keys. It's interesting, and I'm ready to start
    tinkering with my own dictionary for nav & control. I also tried using the
    Georgi's pre-baked Symbol layer, but had some trouble with the alt-tab
    behavior I expect (that I can hold one and tap the other, and the firmware
    will send those intermediate actions despite me still holding alt (or cmd)).  
    It's worth re-flashing with the updated qmk from germ, and if that doesn't
    work I can still try my own symbol map and if THAT doesn't work I can fall
    back to using a steno dictionary entries to define "alt-tab once", "alt-tab
    twice" etc.
- 2020-11-17 Maybe I'll be satisfied looking at the [progress.json activity page](https://github.com/spilliams/steno/commits/main/progress.json)
    for an indication of how consistently I'm practicing. It's no heatcalendar,
    but it'll do.  
    In other news, I've got a cheat sheet of punctuation and whitespacing (the
    book didn't mention, but it's `T*AB` for a tab character). Next item todo
    is to figure out how people manage modifier keys and navigation keys.  
    The real crux is this use case: hold down ctrl (or cmd). Press tab any
    number of times. Also hold shift. Press tab some more times. Let go of
    shift... let go of ctrl (cmd).
- 2020-11-15 Upgraded to Plover v4.0.0. When they say "back up your plover.cfg"
    that file is in `C:\Users\<username>\AppData\Local\plover\plover\`.
    Wondering if there's a way to stay accountable for daily activity without
    bloating this journal file.
- 2020-11-14 journaling with steno now! ok back to qwerty. that took a few
    minutes, but I'm having fun learning about the control mechanisms (to
    insert a space or not, to capitalize or not), as well as punctuation and
    positioning chords. I'd like to try to find a testimonial about Plover's
    "space placement" option (before or after chord output). I feel like that
    option may (a) have a medium to large impact on the workflow for coding vs
    prose, and (b) be really hard to re-train later.
- 2020-11-13 Got a little practice in. I realized I'm not yet ready to fully
    drop the progress-tracking. My reasoning is that I'm doing this from at
    least 2 different computers still, and Typey Type doesn't have user
    accounts, so if I want to track progress there I should copy the file down
    after every session.  
    Thankfully I was able to quickly build a CLI tool that will help me merge
    two JSON files into one, summing the values along the way. That should
    help in the task of saving the progress JSON at least.  
    I also remembered that because Georgi only has the two rows for my fingers
    instead of a steno machine's standard 3, the third thumb button (on both
    sides) is my number bar modifier.  
    I wonder again how I might use the extra two pinky keys on the left side
    too, since I probably won't need--wait...i'm not so terrible at georgi qwerty! [that last phrase only took several minutes to puzzle out, and
    even then i didn't know all the punctuation keys]. Anyway, even while I'm
    doing steno I'll want a way to hit arbitrary shortcuts. That'll be
    interesting.
- 2020-11-12 Started this repo! I hope in the end it'll contain useful tidbits
    I picked up along the journey, as well as any custom dictionaries I develop.
    Some historical background on the journey so far is below. Going forward,
    I don't want this journal to be just a list of wpm counts and times (though
    I'll certainly include some for milestone tracking). The last few times I
    dusted off the steno machine I ended up putting a lot of effort into time-
    and wpm-tracking, so much that it became Not Very Fun Any More. So less of
    that from now on, focussing on the fun parts instead!
- 2020-02-20 [historical] 1 day of practice in February. 13wpm on Georgi.
- 2019-12-16 [historical] 1 day of practice in December. Still hovering around
    10 wpm on Georgi (using Typey Type). I started punctuation though!
- 2019-11-25 [historical] 4 days of practice in October. 1 in November.
- 2019-10-06 [historical] That stint in August lasted 5 days. Seeing a pattern...
    It's ok though. I'm having fun with this when I have fun. Any practical
    effects will be a nice bonus, but not worth beating myself up about if I
    don't achieve.  
    In August I got up to about 16 wpm on Typey Type with the SOFT/HRUF.  
    Now I'm using the Georgi, around 13 wpm on Typey Type.
- 2019-08-25 [historical] That stint in May lasted 4 days. Back at it now. 10
    wpm on Typey Type with the SOFT/HRUF.
- 2019-05-12 [historical] I took baseline typing speed tests at
    www.typingtest.com with my 2015 macbook air and WASD v2 keyboards: 73 and
    76 wpm, respectively.  
    I also started learning the SOFT/HRUF machine, with an initial count of 8
    wpm on Typey Type.
