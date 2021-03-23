- 2921-03-23: Last night and this morning have been trying to get steno working
    on my Ergodox EZ. First milestone achieved was recreating the keymap and
    colors from Oryx. Now it's time to get a steno layer on it! Some early
    trouble:
    - using the existing steno keymap doesn't work with Plover on my work pc.
    - I copied the existing steno map into my new spilliams keymap. This also
        doesn't work but it does make the other layers more friendly.
    - I can already tell I'll have a hard time doing steno with these springs
        (kailh copper I think? 40gf nominal, but tactile makes them 50gf), and
        these keycaps (DSA profile)
    - the georgi does work on the work machine, so it isn't an OS issue with
        COM ports.
    I'll have to come back to this tonight, but some promising early results I
    suppose. I'll have to look into new switches and caps?? I like my latest set
    thouh...
    I can try out the kailh silver (40gf linear), that'll get closer to the
    experience I have with the Georgi (Matias linear reds, 35cN~=35gf), without
    sacrificing the qwerty/gaming performance.
- 2020-11-25 cont'd: I've started a tiny bit of coding with steno. It's alive!
    Starting to think about building my own "single stroke commands" dictionary,
    but because it's so many different chords, thinking of building a generator
    that takes an input of all the chords for each modifier (plus a few others
    for nav/cursor control), and then it prints out all the hundreds of
    permutations! Some things I'm learning about dictionary entries as I do
    this:
    - a `/` is present to delineate strokes
    - a `-` is present to delineate left-hand vs right-hand, but only
        1. if there are no left-hand keys in the chord
        2. if writing it down without a `-` would show ambiguity (`HRPB` is
        either `HR-PB` "license", or `H-RPB` "hit-and-run")
        3. if there is no ambiguity? (e.g. `*ET/K-L` "ethical")
        4. only if there's also no `*`.
        
        Does _Learn Plover!_ have anything to say on this subject?
        > When defining dictionary definitions by hand, you should be sure to
        > include the hyphen when appropriate
        
        I think I can handle that. Why not always generate a `-` until I see it
        break?
    One final note, someone's blog recommended keeping the dictionary handy,
    that being able to quickly add or edit definitions was really helpful. I
    wonder if I could set up a single key in my layer or a brief in a
    dictionary to run that Plover function?
- 2020-11-25 I just realized: using alt to mouse-select multiple cursors in VS
    Code does _not_ trigger the OSL to fall back. Is that a task I do with the
    dictionary instead? No, dictionary can't help there. HMM!  
    Perhaps a key on the symbol layer that does a `TO(BASE)`. yup, that works!  
    I've also printed out some cheat-sheets, and now I guess nothing is
    stopping me from developing some muscle memory...
- 2020-11-24 Made a cheat-sheet for the single-stroke commands dictionary
    tonight. Did my first cut, paste and undo! The shapes are a little goofy
    but I think they'll suffice for my needs. Might even allow me to free up
    some keys on the firmware layer (no need for arrows, letters, or functions).  
    I did order an Ergodox. The clutter argument was very compelling. It'll
    take 4 weeks to get here, so in the meantime I'll keep practicing on Georgi.  
    I'm excited to put together a firmware for Ergodox that does everything I
    need--and with LED control!
- 2020-11-23 A few things I noticed in practice tonight
    - tapping the OSL key doesn't let me alt-tab with impunity, but holding it
        does.
    - my custom nav layer doesn't support select-all, cut, copy, paste, undo,
        or redo. I wonder again if I should record the number of times I use
        a feature in a day, and put those at the front and center.  
        oh, also ctrl-backspace (or alt-bsp and cmd-bsp on mac).  
        Even if I put my most common commands on the first nav layer, I still
        need either  

    1. (hw, current status) a qwerty keyboard standing by  
        pro: I have it (it's free)  
        con: slow to switch from typing to commanding  
        con: BIG desk clutter  
        con: need the qwerty keyboard for typing while learning steno
    2. (hw) a macropad (e.g. Butterstick) standing by  
        pro: I have it (it's free)  
        con: slow to switch from typing to commanding  
        con: smol desk clutter  
        con: need the qwerty keyboard for typing while learning steno
    3. (fw) three or more layers on the georgi  
        pro: I have it (it's free)  
        pro: everything is at my fingertips (it's quick to switch)  
        pro: no clutter!  
        con: patience and time to tweak and learn the layers  
        con: need the qwerty keyboard for typing while learning steno
    4. (hw) A BIGGER georgi, with more keys to hand (see: Gergo. I
        think maybe Gergoplex isn't enough columns)  
        pro: it's a kit I can build myself (another round of custom
        switches, this time maybe kailh silver?)  
        pro: everything is at my fingertips (quick to switch)  
        pro: no clutter!  
        pro: less need to layer the commands (quickER than switching)  
        pro: do not need a qwerty kicking around while I learn steno  
        con: hard to program it to be steno sometimes and command
        sometimes?  
        con: I'd need to build my own enclosure for it.
    5. (hw) white tiger option: get an ergodox already (~$350)  
        pro: it's a kit I can configure myself (it does hotswap)  
        pro: everything is at my fingertips  
        pro: no clutter!  
        pro: less need to layer the commands  
        pro: do not need a qwerty kicking around while I learn steno  
        con: hard to program it to be steno sometimes and command
        sometimes?  
        con: expensive
    6. (sw) learn and develop the Single Stroke commands dictionary in
        Plover  
        pro: I have it (it's free)  
        pro: VERY easy to reconfigure a broken brief on the fly  
        con: need the qwerty keyboard for typing while learning steno  
        con: patience and time to tweak and learn the briefs

    Options 4 and 5 (Gergo and Ergodox) are similar enough that I can
    runoff them: With Gergo, I have to design and build the enclosure,
    and buy keycaps and switches. With Ergodox, I just have to spend
    ~$280 marginally (and maybe make or buy steno-friendly keycaps).
    With Gergo, I get to make it look however I want, with Ergodox, no
    such personalization.  
    ```
    GERGO           ERGODOX
    + personalized  - cost
    - kit work      - not personalized
    ```
    Actually, now that I think about it, the Gergo wouldn't be a full size
    qwerty, so I'd be training (1) steno, (2) columnar qwerty, and (3) gergo
    commands and shortcuts for qwerty. Ergodox loses on customization, but
    "costs" less in terms of money-work, and wins on me not doing that (3)
    training.

    I'd also point out that the "do not need the qwerty keyboard for typing
    while learning steno" is not worth nothing.

    So it's really down to an ideological choice: do I want to do this
    re-training of my brain and fingers through software, firmware, or
    hardware? And even if I do it through software, do I want the greater
    hardware anyway to speed things along, declutter the process, and/or be
    more ergonomic for me? (For the record I have yet to experience any RSI...)
    
    I will sleep on it.

- 2020-11-22 Last night / this morning I read a lot about how QMK works with
    layers. Then I worked up keymaps for both the Butterstick and the Georgi!  
    The first milestone is complete: I can successfully alt-tab any number of
    times!  
    Next milestones:
    - make the nav layer key (and fn layer key) momentary-layers (like OSL, but
    that didn't seem to work immediately).
        > `OSL(layer)` - momentarily activates `layer` until the next key is
        pressed. See One Shot Keys for details and additional functionality.
        ```
        // One-shot layer - 256 layer max
        #define OSL(layer) (QK_ONE_SHOT_LAYER | ((layer)&0xFF))
        ```
        Defined in `quantum/quantum_keycods.h` and used in the `action_for_keycode`
        function of `quantum/keymap_common.c`.
        ```
        #ifndef NO_ACTION_ONESHOT
            case QK_ONE_SHOT_LAYER ... QK_ONE_SHOT_LAYER_MAX:;
                // OSL(action_layer) - One-shot action_layer
                action_layer = keycode & 0xFF;
                action.code  = ACTION_LAYER_ONESHOT(action_layer);
                break;
            case QK_ONE_SHOT_MOD ... QK_ONE_SHOT_MOD_MAX:;
                // OSM(mod) - One-shot mod
                mod         = mod_config(keycode & 0xFF);
                action.code = ACTION_MODS_ONESHOT(mod);
                break;
        #endif
        ```
        which leads to
        ```
        #define ACTION_LAYER_ONESHOT(layer) ACTION_LAYER_TAP((layer), OP_ONESHOT)
        ...
        #define ACTION_LAYER_TAP(layer, key) ACTION(ACT_LAYER_TAP, (layer) << 8 | (key))
        ...
        num action_kind_id { ACT_LAYER_TAP = 0b1010 /* Layer 0-15 */ }
        #define ACTION(kind, param) ((kind) << 12 | (param))
        ```
        So of course all this code is just a bunch of bit shifting.  
        Let's try `OSL()` one more time, just in case...  
        Yeah, it's probably because I had `#define NO_ACTION_ONESHOT` in my
        `config.h`, huh?  
        Nope, `OSL` still isn't working on the georgi. Switching to Butterstick
        to see if it's a steno mode thing.  
        Yes, `OSL` works on the butterstick. Also, implementing
        `oneshot_layer_changed_user` in `keymap.c` makes it print to the QMK
        toolbox (I wasn't seeing those prints from Georgi).  
        Switching back to Georgi for one more shot, then I'll compromise with
        `TG` and hope I don't get in a bad layer state.  
          
        I got it! I had to delete the `#define NO_ACTION_ONESHOT`, _and_ go
        into `rules.mk` (beware there's one for the keyboard and one for the
        keymap!) and set `NO_TAPPING = no`! Now I can use `OSL` to shift layers.  
    - maybe write a custom function like `TH(kc1, kc2)`, where if you tap it
    sends `kc1`, and if you hold it sends `kc2`. For inspiration:
        > `LT(layer, kc)` - momentarily activates `layer` when held, and sends
        > `kc` when tapped. Only supports layers 0-15.
        ```
        // L-ayer, T-ap - 256 keycode max, 16 layer max
        #define LT(layer, kc) (QK_LAYER_TAP | (((layer)&0xF) << 8) | ((kc)&0xFF))
        ...
        // M-od, T-ap - 256 keycode max
        #define MT(mod, kc) (QK_MOD_TAP | (((mod)&0x1F) << 8) | ((kc)&0xFF))
        ```
        Since stubbing this out and then playing with my layers quite a bit
        above, I realized I had an extra blank key next to the media keys, so
        I no longer need to double up `TH(KC_VOLD, KC_MUTE)`, but maybe it'd be
        a fun experiment sometime anyway. Another day. Today my nav and media
        layers work great and I should take some time to practice them (and
        practice steno).
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
    too, since I probably won't need--wait...i'm not so terrible at georgi
    qwerty! [that last phrase only took several minutes to puzzle out, and
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
