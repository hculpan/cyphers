package types

var SubtleCyphers map[string]Cypher = map[string]Cypher{
	"01-04": Cypher{Name: "Analeptic", Level: "1d6 + 2", Effect: `
Restores $(LEVEL) to the user’s Speed Pool.`},
	"05-07": Cypher{Name: "Best tool", Level: "1d6", Effect: `
Provides an additional asset for any one task using a tool, 
even if that means exceeding the normal limit of two assets.`},
	"08-10": Cypher{Name: "Burst of speed", Level: "1d6", Effect: `
For one minute, a user who normally can	move a 
short distance as an action can move a long distance instead.`},
	"11-13": Cypher{Name: "Contingent activator", Level: "1d6 + 2", Effect: `
If the device is activated in conjunction
with another cypher, the user can specify
a condition under which the linked cypher
will activate. The linked cypher retains the
contingent command until it is used (either
normally or contingently). For example,
when this cypher is linked to a cypher that
provides a form of healing or protection,
the user could specify that the linked cypher
will activate if they become damaged to a
certain degree or are subject to a particular
dangerous circumstance. Until the linked
cypher is used, this cypher continues to
count toward the maximum number of
cyphers a PC can carry.`},
	"14-17": Cypher{Name: "Curative", Level: "1d6 + 2", Effect: `
Restores $(LEVEL) to the user’s Might Pool.`},
	"18-20": Cypher{Name: "Darksight", Level: "1d6", Effect: `
Grants the ability to see in the dark for $(LEVEL * 5)
hour(s).`},
	"21-23": Cypher{Name: "Disarm", Level: "1d6 + 1", Effect: `
One NPC within immediate range whose
level is lower than $(LEVEL) drops
whatever they are holding.`},
	"24-26": Cypher{Name: "Eagleseye", Level: "1d6", Effect: `
Grants the ability to see ten times as far
as normal for $(LEVEL) hour(s).`},
	"27-29": Cypher{Name: "Effect resistance", Level: "1d6 + 1", Effect: `
Provides a chance for additional
resistance to directly damaging effects of all
kinds, such as fire, lightning, and the like, for
one day. (It does not provide resistance to
blunt force, slashing, or piercing attacks.) If
the level of the effect is $(LEVEL) or
lower, the user gains an additional defense
roll to avoid it. On a successful defense roll,
treat the attack as if the user had succeeded
on their regular defense roll. (If the user is
an NPC, a PC attacking them with this kind
of effect must succeed on two attack rolls to
harm them.)`},
	"30-32": Cypher{Name: "Effort enhancer (combat)", Level: "1d6 + 1", Effect: `
For the next hour, the user can apply one
free level of Effort to any task (including a
combat task) without spending points from a
Pool. The free level of Effort provided by this
cypher does not count toward the maximum
amount of Effort a character can normally
apply to one task. Once this free level of Effort
is used, the effect of the cypher ends.`},
	"33-35": Cypher{Name: "Effort enhancer (noncombat)", Level: "1d6", Effect: `
For the next hour, the user can apply
one free level of Effort to a noncombat task
without spending points from a Pool. The
level of Effort provided by this cypher does
not count toward the maximum amount of
Effort a character can normally apply to one
task. Once this free level of Effort is used, the
effect of the cypher ends.`},
	"36-39": Cypher{Name: "Enduring shield", Level: "1d6 + 4", Effect: `
For the next day, the user has an asset to
Speed defense rolls.`},
	"40-42": Cypher{Name: "Intellect booster", Level: "1d6 + 2", Effect: `
Adds 1 to the user’s Intellect Edge for one
hour (or 2 if the cypher is level 5 or higher).`},
	"43-45": Cypher{Name: "Intelligence enhancement", Level: "1d6", Effect: `
All of the user’s tasks involving intelligent
deduction—such as playing chess, inferring
a connection between clues, solving a
mathematical problem, finding a bug in
computer code, and so on—are eased by
two steps for one hour. In the subsequent
hour, the strain hinders the same tasks by
two steps.`},
	"46-48": Cypher{Name: "Knowledge enhancement", Level: "1d6", Effect: `
For the next day, the character has training
in a predetermined skill (or two skills if the
cypher is level 5 or higher). The skill could
be anything (including something specific to
the operation of a particular device), or roll a
d100 to choose a common skill.
	01–10 Melee attacks
	11–20 Ranged attacks
	21–40 One type of academic or esoteric lore
		(biology, history, magic, and so on)
	41–50 Repairing (sometimes specific to one
		device)
	51–60 Crafting (usually specific to one
		thing)
	61–70 Persuasion
	71–75 Healing
	76–80 Speed defense
	81–85 Intellect defense
	86–90 Swimming
	91–95 Riding
	96–00 Sneaking`},
	"49-51": Cypher{Name: "Meditation aid", Level: "1d6 + 2", Effect: `
Restores $(LEVEL) to the user’s Intellect Pool.`},
	"52-54": Cypher{Name: "Mind stabilizer", Level: "1d6", Effect: `
The user gains +5 to Armor against
Intellect damage.`},
	"55-57": Cypher{Name: "Motion sensor", Level: "1d6 + 2", Effect: `
For $(LEVEL) hour(s), the user
knows when any movement occurs within
short range, and when large creatures
or objects move within long range (the
cypher distinguishes between the two). It
also indicates the number and size of the
creatures or objects in motion.`},
	"58-60": Cypher{Name: "Nutrition and hydration", Level: "1d6 + 1", Effect: `
The user can go without food and water
for $(LEVEL) days level without ill effect.`},
	"61-63": Cypher{Name: "Perfect memory", Level: "1d6", Effect: `
Allows the user to mentally record
everything they see for $(LEVEL * 30) seconds
and store the recording permanently in their long-term 
memory. This cypher is useful for watching someone pick
a specific lock, enter a complex code, or do
something else that happens quickly.`},
	"64-66": Cypher{Name: "Perfection", Level: "1d6 + 2", Effect: `
The user treats their next action as if they
had rolled a natural 20.`},
	"67-69": Cypher{Name: "Reflex enhancer", Level: "1d6", Effect: `
All tasks involving manual dexterity—such
as pickpocketing, lockpicking, juggling,
operating on a patient, defusing a bomb, and
so on—are eased by two steps for one hour.`},
	"70-73": Cypher{Name: "Rejuvenator", Level: "1d6 + 2", Effect: `
Restores $(LEVEL) points to one random stat Pool.
01–50 Might Pool
51–75 Speed Pool
76–00 Intellect Pool`},
	"74-76": Cypher{Name: "Remembering", Level: "1d6", Effect: `
Allows the user to recall any one
experience they’ve ever had. The experience
can be no longer than one minute per
cypher level, but the recall is perfect, so (for
example) if they saw someone dial a phone,
they will remember the number.`},
	"77-79": Cypher{Name: "Repel", Level: "1d6 + 1", Effect: `
One NPC within immediate range who is
of a level lower than $(LEVEL) decides to
leave, using their next five rounds to move
away quickly.`},
	"80-82": Cypher{Name: "Secret", Level: "1d6 + 2", Effect: `
The user can ask the GM one question
and get a general answer. The GM assigns
a level to the question, so the more obscure
the answer, the more difficult the task.
Generally, knowledge that a PC could find by
looking somewhere other than their current
location is level 1, and obscure knowledge
of the past is level 7. Gaining knowledge of
the future is level 10, and such knowledge
is always open to interpretation. The cypher
cannot provide an answer to a question
above level $(LEVEL).`},
	"83-85": Cypher{Name: "Skill boost", Level: "1d6", Effect: `
Dramatically but temporarily alters the
user’s mind and body so they can ease one
specific kind of physical action by three
steps. Once activated, this boost can be
used $(LEVEL) times , but only within a 
twenty-four-hour period. The boost takes effect 
each time the action is performed. For example,
a level	3 cypher boosts the first three times that
action is attempted. Roll a d100 to determine
the action.
	01–15 Melee attack
	16–30 Ranged attack
	31–40 Speed defense
	41–50 Might defense
	51–60 Intellect defense
	61–68 Jumping
	69–76 Climbing
	77–84 Running
	85–92 Swimming
	93–94 Sneaking
	95–96 Balancing
	97–98 Perceiving
	99 Carrying
	00 Escaping`},
	"86-88": Cypher{Name: "Speed boost", Level: "1d6 + 2", Effect: `
Adds 1 to the user’s Speed Edge for
one hour (adds 2 if the cypher is level 5 or
higher).`},
	"89-91": Cypher{Name: "Stim", Level: "1d6", Effect: `
Eases the user’s next action taken by three steps.`},
	"92-94": Cypher{Name: "Strength boost", Level: "1d6 + 2", Effect: `
Adds 1 to Might Edge for one hour (or 2 if
the cypher is level 5 or higher).`},
	"95-97": Cypher{Name: "Strength enhancer", Level: "1d6", Effect: `
All noncombat tasks involving raw
strength—such as breaking down a door,
lifting a heavy boulder, forcing open
elevator doors, competing in a weightlifting
competition, and so on—are eased by two
steps for one hour.`},
	"98-100": Cypher{Name: "Tissue regeneration", Level: "1d6 + 4", Effect: `
For the next hour, the user regains 1 point
lost to damage per round, up to a total
number of points equal to $(LEVEL * 2). As each point
is regained, they choose which Pool to add it to. 
If all their Pools are at maximum, the regeneration
pauses until they take more damage, at which point
it begins again (if any time remains in the
hour) until the duration expires.`},
}
