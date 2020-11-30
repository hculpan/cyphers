package types

// cypher-data\Numenera.txt

var NumeneraCyphers []Cypher = []Cypher {
	Cypher{ 
		Name: "Adhesion Clamps",
		Level: "1d6",
		Type: []string{
			"Type: Wearable - Gloves",
			"Usable: Handles with powerful suction cups",
		},
		Effect: `Allows for automatic climbing of any
surface, even horizontal ones. Lasts for ten
minutes per cypher level.`,
	},
	Cypher{ 
		Name: "Antivenom",
		Level: "1d6 + 2",
		Type: []string{
			"Internal: Pill, ingestible liquid",
			"Usable: Injector",
		},
		Effect: `Renders user immune to poisons of the
same level or lower for one hour per cypher
level and ends any such ongoing effects, if
any, already in the user’s system.`,
	},
	Cypher{ 
		Name: "Attractor",
		Level: "1d6 + 4",
		Type: []string{
			"Type: Wearable - Glove of synth",
			"Usable: Small handheld device",
		},
		Effect: `One unanchored item your size or smaller
within long range (very long range if the cypher
is level 8 or higher) is drawn immediately to
the device. This takes one round. The item has
no momentum when it arrives.`,
	},
	Cypher{ 
		Name: "Banishing Nodule",
		Level: "1d6",
		Type: []string{
			"Usable: Crystal nodule affixed to a melee",
		},
		Effect: `For the next 28 hours, each time the
weapon the nodule is attached to strikes a
solid creature or object, it generates a burst
of energy that teleports the creature or object
struck an immediate distance in a random
direction (not up or down). The teleported
creature’s actions (including defense) are
hindered on its next turn (hindered by two
steps if the cypher is level 5 or higher).`,
	},
	Cypher{ 
		Name: "Blinking Nodule",
		Level: "1d6",
		Type: []string{
			"Usable: Crystal nodule affixed to armor",
		},
		Effect: `For the next 28 hours, each time (but
not more than once per round) the wearer
of the armor the nodule is attached to is
struck hard enough to inflict damage, they
teleport an immediate distance in a random
direction (not up or down). Since the wearer
is prepared for this effect and their foe is
not, the wearer’s defenses are eased for
one round after they teleport (eased by two
steps if the cypher is level 5 or higher).`,
	},
	Cypher{ 
		Name: "Catholicon",
		Level: "1d6 + 2",
		Type: []string{
			"Internal: Pill, ingestible liquid",
			"Usable: Injector",
		},
		Effect: `Cures any disease of the cypher level
or lower.`,
	},
	Cypher{ 
		Name: "Catseye",
		Level: "1d6",
		Type: []string{
			"Internal: Pill, ingestible liquid",
			"Usable: Injector",
		},
		Effect: `Grants the ability to see in the dark for
five hours per cypher level.`,
	},
	Cypher{ 
		Name: "Chemical Factory",
		Level: "1d6",
		Type: []string{
			"Internal: Pill, ingestible liquid",
			"Usable: Injector",
		},
		Effect: `After one hour, the sweat of the user
produces 1d6 doses of a valuable liquid (these
doses are not considered cyphers). They must
be used within one week. Effects vary:
01–04 Euphoric for 1d6 hours
05–08 Hallucinogenic for 1d6 hours
09–12 Stimulant for 1d6 hours
13–16 Depressant for 1d6 hours
17–20 Nutrient supplement
21–25 Antivenom
26–30 Cures disease
31–35 See in the dark for one hour
36–45 Restores a number of Might Pool
points equal to cypher level
46–55 Restores a number of Speed Pool
points equal to cypher level
56–65 Restores a number of Intellect Pool
points equal to cypher level
66–75 Increases Might Edge by 1 for one
hour
76–85 Increases Speed Edge by 1 for one
hour
86–95 Increases Intellect Edge by 1 for one
hour
96–00 Restores all Pools to full`,
	},
	Cypher{ 
		Name: "Comprehension Graft",
		Level: "1d6 + 1",
		Type: []string{
			"Usable: Small metallic disk",
		},
		Effect: `When applied to a creature’s head, the
disk immediately unleashes microfilaments
that enter the brain. Within five minutes,
the creature can understand the words of
a specific language keyed to the graft (two
languages if the cypher is level 5 or higher).
This is true even of creatures that do not
normally have a language. If the creature
could already understand the language,
the cypher has no effect. Once the graft
attaches, the effect is permanent, and
this device no longer counts against the
number of cyphers that a PC can bear.`,
	},
	Cypher{ 
		Name: "Controlled Blinking Nodule",
		Level: "1d6 + 2",
		Type: []string{
			"Usable: Crystal nodule affixed to armor",
		},
		Effect: `For the next 28 hours, each time the
wearer of the armor the nodule is attached to
is struck hard enough to inflict damage (but
no more than once per round), they teleport
to a spot they desire within immediate
range. Since the wearer is prepared for
this effect and their foe is not, the wearer’s
defenses are eased for one round after they
teleport (eased by two steps if the cypher is
level 6 or higher).`,
	},
	Cypher{ 
		Name: "Datasphere Siphon",
		Level: "1d6",
		Type: []string{
			"Internal: Pill, ingestible liquid",
			"Type: Wearable - Temporary tattoo, amulet,",
			"Usable: Small handheld device, crystal",
		},
		Effect: `Tapping into the datasphere’s
knowledge, the user can learn the answer
to one question (two questions if the
cypher is level 4 or higher, three questions
if the cypher is level 6 or higher).`,
	},
	Cypher{ 
		Name: "Density Nodule",
		Level: "1d6",
		Type: []string{
			"Usable: Crystal nodule affixed to a melee",
		},
		Effect: `For the next 28 hours, each time the
weapon the nodule is attached to strikes
a solid creature or object, the weapon
suddenly increases dramatically in weight,
causing the blow to inflict an additional 2
points of damage (3 points if the cypher is
level 4 or higher).`,
	},
	Cypher{ 
		Name: "Detonation",
		Level: "1d6 + 2",
		Type: []string{
			"Type: Wearable - Wristband projector (long range)",
			"Usable: Explosive device (thrown, short range)",
		},
		Effect: `Explodes in an immediate radius,
inflicting damage equal to the cypher level.
Roll for the type of damage:
01–10 Cell-disrupting (harms only flesh)
11–30 Corrosive
31–40 Electrical discharge
41–50 Heat drain (cold)
51–75 Fire
76–00 Shrapnel`,
	},
	Cypher{ 
		Name: "Detonation (desiccating)",
		Level: "1d6 + 2",
		Type: []string{
			"Type: Wearable - Wristband projector (long range)",
			"Usable: Explosive device or ceramic sphere",
		},
		Effect: `Bursts in an immediate radius, draining
moisture from everything within it. Living
creatures take damage equal to the cypher
level. Water in the area is vaporized.`,
	},
	Cypher{ 
		Name: "Detonation (flash)",
		Level: "1d6 + 2",
		Type: []string{
			"Type: Wearable - Wristband projector (long range)",
			"Usable: Explosive device or ceramic sphere",
		},
		Effect: `Bursts in an immediate radius, blinding
all within it for one minute (ten minutes if
the cypher is level 4 or higher).`,
	},
	Cypher{ 
		Name: "Detonation (gravity)",
		Level: "1d6 + 2",
		Type: []string{
			"Type: Wearable - Wristband projector (long range)",
			"Usable: Explosive device or ceramic sphere",
		},
		Effect: `Bursts in an immediate radius, inflicting
damage equal to the cypher level by
increasing gravity tremendously for one
second. All in the area are crushed to the
ground for one round and cannot take
physical actions.`,
	},
	Cypher{ 
		Name: "Detonation (massive)",
		Level: "1d6 + 2",
		Type: []string{
			"Type: Wearable - Wristband projector (very long range)",
			"Usable: Handheld projector (very long range)",
		},
		Effect: `Explodes in a short-range radius,
inflicting damage equal to the cypher level.
Roll for the type of damage:
01–10 Cell-disrupting (harms only flesh)
11–30 Corrosive
31–40 Electrical discharge
41–50 Heat drain (cold)
51–75 Fire
76–00 Shrapnel`,
	},
	Cypher{ 
		Name: "Detonation (matter Disruption)",
		Level: "1d6 + 4",
		Type: []string{
			"Type: Wearable - Wristband projector (long range)",
			"Usable: Explosive device or ceramic sphere",
		},
		Effect: `Explodes in an immediate radius,
releasing nanites that rearrange matter in
random ways. Inflicts damage equal to the
cypher level.`,
	},
	Cypher{ 
		Name: "Detonation (pressure)",
		Level: "1d6 + 2",
		Type: []string{
			"Type: Wearable - Wristband projector (long range)",
			"Usable: Explosive device (thrown, short range)",
		},
		Effect: `Explodes in an immediate radius,
inflicting impact damage equal to the
cypher level. Also moves unattended
objects out of the area if they weigh less
than 20 pounds (9 kg) per cypher level.`,
	},
	Cypher{ 
		Name: "Detonation (singularity)",
		Level: "10",
		Type: []string{
			"Usable: Explosive device or ceramic sphere",
		},
		Effect: `Explodes and creates a momentary
singularity that tears at the fabric of the
universe. Inflicts 20 points of damage to all
within short range, drawing them (or their
remains) together to immediate range (if
possible). Player characters in the radius
move one step down the damage track if
they fail a Might defense roll.`,
	},
	Cypher{ 
		Name: "Detonation (sonic)",
		Level: "1d6 + 2",
		Type: []string{
			"Type: Wearable - Wristband projector (long range)",
			"Usable: Explosive device or ceramic sphere",
		},
		Effect: `Explodes with terrifying sound,
deafening all in an immediate radius for
ten minutes per cypher level.`,
	},
	Cypher{ 
		Name: "Detonation (spawn)",
		Level: "1d6 + 2",
		Type: []string{
			"Type: Wearable - Wristband projector (long range)",
			"Usable: Explosive device or ceramic sphere",
		},
		Effect: `Bursts in an immediate radius, blinding
all within it for one minute and inflicting
damage equal to the cypher level. The
burst spawns 1d6 additional detonations;
on the next round, each additional
detonation flies to a random spot within
short range and explodes in an immediate
radius. Roll for the type of damage dealt by
all detonations:
01–10 Cell-disrupting (harms only flesh)
11–30 Corrosive
31–40 Electrical discharge
41–50 Heat drain (cold)
51–75 Fire
76–00 Shrapnel`,
	},
	Cypher{ 
		Name: "Detonation (web)",
		Level: "1d6 + 2",
		Type: []string{
			"Type: Wearable - Wristband projector (long range)",
			"Usable: Explosive device (thrown, short range)",
		},
		Effect: `Explodes in an immediate radius and
creates sticky strands of goo that last
1 hour. PCs caught in the area must use
a Might-based action to get out, with the
difficulty determined by the cypher level.
NPCs break free if their level is higher than
the cypher level.`,
	},
	Cypher{ 
		Name: "Disrupting Nodule",
		Level: "1d6",
		Type: []string{
			"Usable: Crystal nodule affixed to a melee",
		},
		Effect: `For the next 28 hours, each time the
weapon the nodule is attached to strikes
a solid creature or object, it generates a
burst of nanites that directly attack organic
cells. The affected target takes 1 additional
point of damage (2 points if the cypher is
level 4 or higher, 3 points if the cypher is
level 6 or higher) and loses its next action.`,
	},
	Cypher{ 
		Name: "Eagleseye",
		Level: "1d6",
		Type: []string{
			"Internal: Pill, ingestible liquid",
			"Usable: Injector",
		},
		Effect: `Grants the ability to see ten times as far
as normal for one hour per cypher level.`,
	},
	Cypher{ 
		Name: "Fireproofing Spray",
		Level: "1d6 + 4",
		Type: []string{
			"Usable: Spray canister",
		},
		Effect: `An object sprayed by this cypher has
Armor against fire damage equal to the
cypher’s level for 28 hours.`,
	},
	Cypher{ 
		Name: "Flame-Retardant Wall",
		Level: "1d6",
		Type: []string{
			"Type: Wearable - Belt, ring, bracelet",
			"Usable: Handheld device",
		},
		Effect: `Creates an immobile plane of permeable
energy up to 20 feet by 20 feet (6 m by 6 m)
for one hour per cypher level. The plane
conforms to the space available. Flames
passing through the plane are extinguished.`,
	},
	Cypher{ 
		Name: "Force Cube Projector",
		Level: "1d6 + 3",
		Type: []string{
			"Type: Wearable - Belt, ring, bracelet",
			"Usable: Handheld device",
		},
		Effect: `Creates an immobile cube composed of
six planes of solid force, each 30 feet (9 m)
to a side, for one hour. The planes conform
to the space available.`,
	},
	Cypher{ 
		Name: "Force Nodule",
		Level: "1d6",
		Type: []string{
			"Usable: Crystal nodule affixed to armor",
		},
		Effect: `For the next 28 hours, the armor the nodule
is attached to is bolstered by a powerful force
field, adding 2 to the Armor it provides (adding
3 to the Armor if the cypher is level 5 or higher).`,
	},
	Cypher{ 
		Name: "Force Screen Projector",
		Level: "1d6 + 3",
		Type: []string{
			"Type: Wearable - Belt, ring, bracelet",
			"Usable: Handheld device",
		},
		Effect: `Creates an immobile plane of solid
force up to 20 feet by 20 feet (6 m by 6 m)
for one hour per cypher level. The plane
conforms to the space available.`,
	},
	Cypher{ 
		Name: "Force Shield Projector",
		Level: "1d6 + 3",
		Type: []string{
			"Internal: Subdermal injection",
			"Type: Wearable - Belt, ring, bracelet",
			"Usable: Handheld device",
		},
		Effect: `Creates a shimmering energy shield
around the user for one hour, during which
time they gain +3 Armor (+4 Armor if the
cypher is level 5 or higher).`,
	},
	Cypher{ 
		Name: "Friction-Reducing Gel",
		Level: "1d6",
		Type: []string{
			"Usable: Spray canister",
		},
		Effect: `Sprayed across an area up to 10 feet (3
m) square, this gel makes things extremely
slippery. For one hour per cypher level,
movement tasks in the area are hindered
by three steps.`,
	},
	Cypher{ 
		Name: "Frigid Wall Projector",
		Level: "1d6 + 2",
		Type: []string{
			"Usable: Complex device",
		},
		Effect: `Creates a wall of supercooled air up
to 30 feet by 30 feet by 1 foot (9 m by 9 m
by 30 cm) that inflicts damage equal to
the cypher level on anything that passes
through it. The wall conforms to the space
available. It lasts for ten minutes.`,
	},
	Cypher{ 
		Name: "Gas Bomb",
		Level: "1d6 + 2",
		Type: []string{
			"Type: Wearable - Wristband projector (long range)",
			"Usable: Explosive device or ceramic sphere",
		},
		Effect: `Bursts in a poisonous cloud within an
immediate distance. The cloud lingers
for 1d6 rounds unless conditions dictate
otherwise. Effects vary:
01–10 Thick smoke: occludes sight while
the cloud lasts.
11–20 Choking gas: living creatures that
breathe lose their actions to choking
and coughing for a number of
rounds equal to the cypher level.
21–50 Poison gas: living creatures that
breathe suffer damage equal to the
cypher level.
51–60 Corrosive gas: everything suffers
damage equal to the cypher level.
61–65 Hallucinogenic gas: living creatures
that breathe lose their actions to
hallucinations and visions for a number
of rounds equal to the cypher level.
66–70 Nerve gas: living creatures that
breathe suffer Speed damage equal
to the cypher level.
71–80 Mind-numbing gas: living creatures
that breathe suffer Intellect damage
equal to the cypher level.
81–83 Fear gas: living creatures that
breathe and think flee in a random
direction in fear (or are paralyzed
with fear) for a number of rounds
equal to the cypher level.
84–86 Amnesia gas: living creatures that
breathe and think permanently lose
all memory of the last minute.
87–96 Sleep gas: living creatures that
breathe fall asleep for a number of
rounds equal to the cypher level or
until awoken by a violent action or
an extremely loud noise.
97–00 Rage gas: living creatures that
breathe and think make a melee
attack on the nearest creature and
continue to do so for a number of
rounds equal to the cypher level.`,
	},
	Cypher{ 
		Name: "Gravity Nullifier",
		Level: "1d6 + 3",
		Type: []string{
			"Internal: Subdermal injection",
			"Type: Wearable - Belt, boots, ring, bracelet",
			"Usable: Small platform on which the user",
		},
		Effect: `For one hour, the user can float into the
air, moving vertically (but not horizontally
without some other action, such as pushing
along the ceiling) up to a short distance per
round. The user must weigh less than 50
pounds (22 kg) per level of the cypher.`,
	},
	Cypher{ 
		Name: "Gravity-Nullifying Spray",
		Level: "1d6 + 2",
		Type: []string{
			"Usable: Spray canister",
		},
		Effect: `A nonliving object up to the size of a human
(two humans if the cypher is level 6 or higher)
sprayed by this cypher floats 1d20 feet in the
air permanently and no longer has weight if
carried, though it needs to be strapped down.`,
	},
	Cypher{ 
		Name: "Heat Nodule",
		Level: "1d6",
		Type: []string{
			"Usable: Crystal nodule affixed to a melee weapon",
		},
		Effect: `For the next 28 hours, each time the weapon
the nodule is attached to strikes a solid
creature or object, it generates a burst of heat,
inflicting an additional 2 points of damage
(3 points if the cypher is level 4 or higher, 4
points if the cypher is level 6 or higher).`,
	},
	Cypher{ 
		Name: "Hunter/seeker",
		Level: "1d6",
		Type: []string{
			"Type: Wearable - Arm- or shoulder-mounted launcher",
			"Usable: Complex device, handheld device",
		},
		Effect: `With long-range movement, this intelligent
missile tracks and attacks a specified target
(target must be within sight when selected). If
it misses, it continues to attack one additional
time per cypher level until it hits. For example,
a level 4 hunter/seeker will attack a maximum
of five times. Different hunter/seekers have
different effects:
01–50 Inflicts 8 points of damage.
51–80 Bears a poisoned needle that inflicts
3 points of damage plus poison.
81–90 Explodes, inflicting 6 points of
damage to all within immediate range.
91–95 Shocks for 4 points of electricity
damage, and stuns for one round
per cypher level.
96–00 Covers target in sticky goo that
immediately hardens, holding them
fast until they break out with a
Might action (difficulty equal to the
cypher level + 2).`,
	},
	Cypher{ 
		Name: "Image Projector",
		Level: "1d6",
		Type: []string{
			"Type: Wearable - Headband with device on forehead",
			"Usable: Handheld device with glass panel",
		},
		Effect: `Projects one of the following immobile
images in the area described for one hour.
The image appears up to a close distance
away (long distance if the cypher level
is 4 or higher, very long distance if the
cypher level is 6 or higher). Scenes include
movement, sound, and smell.
01–20 Terrifying creature of an unknown
species, perhaps no longer alive in
the world (10-foot [3 m] cube)
21–40 Huge machine that obscures sight
(30-foot [9 m] cube)
41–50 Beautiful pastoral scene (50-foot
[15 m] cube)
51–60 Food that looks delicious but may
not be familiar (10-foot [3 m] cube)
61–80 Solid color that obscures sight
(50-foot [15 m] cube)
81–00 Incomprehensible scene that is
disorienting and strange (20-foot
[6 m] cube)`,
	},
	Cypher{ 
		Name: "Inferno Wall Projector",
		Level: "1d6 + 2",
		Type: []string{
			"Usable: Complex device",
		},
		Effect: `Creates a wall of extreme heat up to
30 feet by 30 feet by 1 foot (9 m by 9 m
by 30 cm) that inflicts damage equal to
the cypher level on anything that passes
through it. The wall conforms to the space
available. It lasts for ten minutes.`,
	},
	Cypher{ 
		Name: "Infiltrator",
		Level: "1d6",
		Type: []string{
			"Internal: Phases into eye, phases out when",
			"Type: Wearable - Adheres to temple and launches",
			"Usable: Handheld device that launches",
		},
		Effect: `Tiny capsule launches and moves at
great speed, mapping and scanning an
unknown area. It moves 500 feet (150 m)
per level, scanning an area up to 50 feet
(15 m) per level away from it. It identifies
basic layout, creatures, and major energy
sources. Its movement is blocked by any
physical or energy barrier.`,
	},
	Cypher{ 
		Name: "Instant Servant",
		Level: "1d6",
		Type: []string{
			"Usable: Handheld device",
		},
		Effect: `Small device expands into a humanoid
automaton that is roughly 2 feet (60 cm)
tall. Its level is equal to the cypher level and
it can understand the verbal commands
of the character who activates it. Once
the servant is activated, commanding it
is not an action. It can make attacks or
perform actions as ordered to the best of
its abilities, but it cannot speak.
The automaton has short-range movement
but never goes farther than long range
away from the character who activated it.
At the GM’s discretion, the servant might
have specialized knowledge, such as how
to operate a particular device. Otherwise,
it has no special knowledge. In any case,
the servant is not artificially intelligent or
capable of initiating action. It does only as
commanded.
The servant operates for one hour per cypher
level.`,
	},
	Cypher{ 
		Name: "Instant Shelter",
		Level: "1d6 + 3",
		Type: []string{
			"Usable: Handheld device",
		},
		Effect: `With the addition of water and air,
the small device expands into a simple
one-room structure with a door and a
transparent window (two rooms with
an internal door if the cypher is level 7
or higher). The structure is 10 feet by 10
feet by 20 feet (3 m by 3 m by 6 m). It is
made from a form of shapestone and is
permanent and immobile once created.`,
	},
	Cypher{ 
		Name: "Intellect Enhancement",
		Level: "1d6 + 2",
		Type: []string{
			"Internal: Pill, ingestible liquid",
			"Type: Wearable - Adhesive patch that activates when",
			"Usable: Injector",
		},
		Effect: `Substance adds 1 to Intellect Edge for
one hour (or adds 2 if the cypher is level 5
or higher).`,
	},
	Cypher{ 
		Name: "Invisibility Nodule",
		Level: "1d6",
		Type: []string{
			"Usable: Crystal nodule affixed to armor",
		},
		Effect: `For the next ten hours per cypher level,
the armor the nodule is attached to is
invisible, making the wearer appear to be
unarmored.`,
	},
	Cypher{ 
		Name: "Knowledge Enhancement",
		Level: "1d6",
		Type: []string{
			"Internal: Pill, ingestible liquid",
			"Type: Wearable - Adhesive patch that activates when",
			"Usable: Injector",
		},
		Effect: `For the next 28 hours, the character has
training in a predetermined skill (or two skills
if the cypher is level 5 or higher). Although the
skill could be anything (including something
specific to the operation of one device or
something similar), common skills include:
01–10 Melee attacks
11–20 Ranged attacks
21–40 Understanding numenera
(sometimes specific to one device)
41–50 Repairing (sometimes specific to
one device)
51–60 Crafting (usually specific to one
thing)
61–70 Persuasion
71–75 Healing
76–80 Speed defense
81–85 Intellect defense
86–90 Swimming
91–95 Riding
96–00 Sneaking`,
	},
	Cypher{ 
		Name: "Lightning Wall Projector",
		Level: "1d6 + 2",
		Type: []string{
			"Usable: Complex device",
		},
		Effect: `Creates a wall of electric bolts up to
30 feet by 30 feet by 1 foot (9 m by 9 m
by 30 cm) that inflicts damage equal to
the cypher level on anything that passes
through it. The wall conforms to the space
available. It lasts for ten minutes.`,
	},
	Cypher{ 
		Name: "Living Solvent",
		Level: "1d10",
		Type: []string{
			"Usable: Canister containing slime",
		},
		Effect: `Once released, this organic slime
dissolves 1 cubic foot of material each
round. After one round per cypher level,
the slime dies and becomes inert.`,
	},
	Cypher{ 
		Name: "Machine Control Implant",
		Level: "1d6 + 2",
		Type: []string{
			"Internal: Pill",
			"Type: Wearable - Disk that adheres to forehead,",
			"Usable: Injector",
		},
		Effect: `When activated, the cypher splits into
two pieces. One is affixed to a numenera
device and the other to a character. The
character can then use their mind to
control the device at long range, bidding it
to do anything it could do normally. Thus,
a device could be activated or deactivated,
and a vehicle could be piloted. The control
lasts for ten minutes per cypher level, and
once the device is chosen, it cannot be
changed.`,
	},
	Cypher{ 
		Name: "Magnetic Attack Drill",
		Level: "1d6 + 2",
		Type: []string{
			"Usable: Small sphere with a thick screw",
		},
		Effect: `The user throws this cypher at a target
within short range, and it drills into the
target for one round, inflicting damage
equal to the cypher level. If the target is
made of metal or wearing metal (such as
armor), the attack is eased.`,
	},
	Cypher{ 
		Name: "Magnetic Master",
		Level: "1d6 + 2",
		Type: []string{
			"Type: Wearable - Gloves with metal plates",
			"Usable: Small pyramid-shaped metallic device",
		},
		Effect: `Establishes a connection with one
metal object within short range that a
human could hold in one hand. After this
connection is established, the user can
move or manipulate the object anywhere
within short range (each movement or
manipulation is an action). For example,
the user could wield a weapon or drag a
helm affixed to a foe’s head to and fro. The
connection lasts for ten rounds per cypher
level.`,
	},
	Cypher{ 
		Name: "Magnetic Shield",
		Level: "1d6 + 2",
		Type: []string{
			"Type: Wearable - Gloves with metal plates",
			"Usable: Small pyramid-shaped metallic device",
		},
		Effect: `For ten minutes per cypher level, metal
objects cannot come within immediate
range of the activated device. Metal items
already in the area when the device is
activated are slowly pushed out.`,
	},
	Cypher{ 
		Name: "Memory Lenses",
		Level: "1d6",
		Type: []string{
			"Type: Wearable - Contact lenses, eyeglasses, or",
		},
		Effect: `Allows the wearer to mentally record
everything they see for thirty seconds
per cypher level and store the recording
permanently in their long-term memory.
This cypher is useful for watching someone
pick a specific lock, enter a complex code,
or do something else that happens quickly.`,
	},
	Cypher{ 
		Name: "Mental Scrambler",
		Level: "1d6 + 2",
		Type: []string{
			"Usable: Complex metal and glass device",
		},
		Effect: `Two rounds after being activated, the
device creates an invisible field that fills an
area within short range and lasts for one
minute. The field scrambles the mental
processes of all thinking creatures. The
effect lasts as long as they remain in the
field and for 1d6 rounds after, although an
Intellect defense roll is allowed each round
to act normally (both in the field and after
leaving it). Each mental scrambler is keyed
to a specific effect. Roll for effect:
01–30 Victims cannot act.
31–40 Victims cannot speak.
41–50 Victims move slowly (immediate
range) and clumsily.
51–60 Victims cannot see or hear.
61–70 Victims lose all sense of direction,
depth, and proportion.
71–80 Victims do not recognize anyone
they know.
81–88 Victims suffer partial amnesia.
89–94 Victims suffer total amnesia.
95–98 Victims lose all inhibitions,
revealing secrets and performing
surprising actions.
99–00 Victims’ ethics are inverted.`,
	},
	Cypher{ 
		Name: "Metal Death",
		Level: "1d6 + 2",
		Type: []string{
			"Type: Wearable - Wrist-mounted sprayer",
			"Usable: Canister with hose",
		},
		Effect: `Produces a stream of foam that covers
an area about 3 feet by 3 feet (1 m by 1 m),
transforming any metal that it touches into
a substance as brittle as thin glass. The
foam affects metal to a depth of about 6
inches (15 cm).`,
	},
	Cypher{ 
		Name: "Monoblade",
		Level: "1d6 + 2",
		Type: []string{
			"Internal: Injection into fingertip",
			"Type: Wearable - Glove",
			"Usable: Device similar to hilt",
		},
		Effect: `Produces a 6-inch (15 cm) blade that’s
the same level as the cypher. The blade cuts
through any material of a level lower than its
own. If used as a weapon, it is a light weapon
that ignores Armor of a level lower than its
own. The blade lasts for ten minutes.`,
	},
	Cypher{ 
		Name: "Motion Sensor",
		Level: "1d6 + 2",
		Type: []string{
			"Internal: Injection into spine",
			"Type: Wearable - Amulet",
			"Usable: Disk that can be affixed to the floor or",
		},
		Effect: `Indicates when any movement occurs
within short range, or when large creatures
or objects move within long range (the
cypher distinguishes between the two).
It also indicates the number and size of
the creatures or objects in motion. Once
activated, it operates for one hour per
cypher level.`,
	},
	Cypher{ 
		Name: "Personal Environment Field",
		Level: "1d6 + 2",
		Type: []string{
			"Type: Wearable - Belt, medallion, ring",
			"Usable: Handheld device",
		},
		Effect: `Creates an aura of temperature and
atmosphere that will sustain a human
safely for 28 hours. The aura extends to 1
foot (30 cm) around the user (double that
radius if the cypher is level 7 or higher). It
does not protect against sudden flashes of
temperature change (such as from a heat
ray). A small number of these cyphers (1%)
accommodate the preferred environment
of a nonhuman, nonterrestrial creature.`,
	},
	Cypher{ 
		Name: "Phase Changer",
		Level: "1d6 + 1",
		Type: []string{
			"Type: Wearable - Belt, medallion, ring",
			"Usable: Handheld device",
		},
		Effect: `Puts the user out of phase for one
minute (two minutes if the cypher is level
6 or higher). During this time, the user can
pass through solid objects as though they
were entirely insubstantial, like a ghost.
They cannot make physical attacks or be
physically attacked.`,
	},
	Cypher{ 
		Name: "Phase Disruptor",
		Level: "1d6 + 2",
		Type: []string{
			"Usable: Complex device, plate that affixes to",
		},
		Effect: `Puts a portion of a physical structure (like
a wall or floor) out of phase for one hour. It
affects an area equal to one 5-foot (1.5 m)
cube per cypher level. While the area is out of
phase, creatures and objects can pass freely
through it as if it were not there, although
one cannot see through it, and it blocks light.`,
	},
	Cypher{ 
		Name: "Poison (emotion)",
		Level: "1d6 + 2",
		Type: []string{
			"Internal: Pill, ingestible or injectable liquid",
			"Type: Wearable - Lipstick, false fingertip, ring with needle",
			"Usable: Injector",
		},
		Effect: `The victim feels a specific emotion for
one hour.
01–20 Anger. Likely to attack anyone who
disagrees with them. Very hard to
interact with; all interaction tasks
are hindered by two steps.
21–40 Fear. Flees in terror for one minute
when threatened.
41–60 Lust. Cannot focus on any
nonsexual activity.
61–75 Sadness. All tasks are hindered.
76–85 Complacency. Has no motivation.
All tasks are hindered by two steps.
86–95 Joy. Easy to interact with in a
pleasant manner; all pleasant
interaction tasks are eased.
96–00 Love. Much easier to interact
with; all interaction tasks are
eased by two steps, but temporary
attachment is likely.`,
	},
	Cypher{ 
		Name: "Poison (explosive)",
		Level: "1d6 + 1",
		Type: []string{
			"Internal: Pill, ingestible or injectable liquid",
			"Type: Wearable - Lipstick, false fingertip, ring with",
			"Usable: Injector",
		},
		Effect: `Once this substance enters the
bloodstream, it travels to the brain
and reorganizes into an explosive that
detonates when activated, inflicting 10
points of damage (ignoring Armor). Roll to
determine the means of detonation:
01–25 The detonator is activated (must be
within long range).
26–40 A specified amount of time passes.
41–50 The victim takes a specific action.
51–55 A specific note is sung or played on
an instrument within short range.
56–60 The victim smells a specific scent
within immediate range.
61–80 The victim comes within long range
of the detonator.
81–00 The victim is no longer within long
range of the detonator.`,
	},
	Cypher{ 
		Name: "Poison (mind-Controlling)",
		Level: "1d6 + 2",
		Type: []string{
			"Internal: Pill, ingestible or injectable liquid",
			"Type: Wearable - Lipstick, false fingertip, ring with",
			"Usable: Injector",
		},
		Effect: `The victim must carry out a specific
action in response to a specific trigger.
01–20 Lies down for one minute with eyes
closed when told to do so.
21–40 Flees in terror for one minute when
threatened.
41–60 Answers questions truthfully for
one minute.
61–75 Attacks close friend for one round
when within immediate range.
76–85 Obeys next verbal command given
(if it is understood).
86–95 For 28 hours, becomes sexually
attracted to the next creature of its
own species that it sees.
96–00 For one minute, moves toward the
next red object seen in lieu of all
other actions, even ignoring selfpreservation.`,
	},
	Cypher{ 
		Name: "Poison (mind-Disrupting)",
		Level: "1d6 + 2",
		Type: []string{
			"Internal: Pill, ingestible or injectable liquid",
			"Type: Wearable - Lipstick, false fingertip, ring with",
			"Usable: Injector",
		},
		Effect: `The victim suffers Intellect damage
equal to the cypher’s level and cannot take
actions for a number of rounds equal to
the cypher’s level.`,
	},
	Cypher{ 
		Name: "Psychic Communique",
		Level: "1d6 + 2",
		Type: []string{
			"Internal: Pill",
			"Type: Wearable - Device that adheres to temple",
			"Usable: Metallic disk",
		},
		Effect: `Allows the user to project a one-time,
one-way telepathic message of up to ten
words per cypher level, with an unlimited
range, to anyone they know.`,
	},
	Cypher{ 
		Name: "Ray Emitter",
		Level: "1d6 + 2",
		Type: []string{
			"Internal: Subdermal implant",
			"Type: Wearable - Contact lens, glove, ring, wristband,",
			"Usable: Handheld device",
		},
		Effect: `Allows the user to project a ray of
destructive energy up to very long range
that inflicts damage equal to the cypher’s
level.
01–50 Heat/concentrated light
51–60 Cell-disrupting radiation
61–80 Force
81–87 Magnetic wave
88–93 Molecular bond disruption
94–00 Concentrated cold`,
	},
	Cypher{ 
		Name: "Ray Emitter (numbing)",
		Level: "1d6 + 2",
		Type: []string{
			"Internal: Subdermal implant",
			"Type: Wearable - Contact lens, glove, ring, wristband,",
			"Usable: Handheld device",
		},
		Effect: `Allows the user to project a ray of
energy up to long range (very long range if
the cypher is level 6 or higher) that numbs
one limb of the target, making it useless
for one minute. A small number of these
devices (5%) induce numbing that lasts for
one hour.`,
	},
	Cypher{ 
		Name: "Ray Emitter (paralysis)",
		Level: "1d6 + 2",
		Type: []string{
			"Internal: Subdermal implant",
			"Type: Wearable - Contact lens, glove, ring,",
			"Usable: Handheld device",
		},
		Effect: `Allows the user to project a ray of
energy up to very long range that paralyzes
the target for one minute. A small number
of these devices (5%) induce paralysis that
lasts for one hour.`,
	},
	Cypher{ 
		Name: "Reality Spike",
		Level: "1d6 + 4",
		Type: []string{
			"Usable: Metallic spike",
		},
		Effect: `Once activated, the spike does not
move—ever—even if activated in midair.
A Might action will dislodge the spike, but
then it is ruined.`,
	},
	Cypher{ 
		Name: "Rejuvenator",
		Weight: 2,
		Level: "1d6 + 2",
		Type: []string{
			"Internal: Pill, ingestible liquid",
			"Type: Wearable - Adhesive patch that activates when",
			"Usable: Injector",
		},
		Effect: `Substance restores a number of points
equal to the cypher’s level to one random
Pool. Roll 1d100:
01–50 Might Pool
51–75 Speed Pool
76–00 Intellect Pool`,
	},
	Cypher{ 
		Name: "Remote Viewer",
		Level: "1d6",
		Type: []string{
			"Usable: Device that splits into two parts when",
		},
		Effect: `For one hour per cypher level, the glass
screen on one part shows everything going
on in the vicinity of the other part, regardless
of the distance between the two parts.`,
	},
	Cypher{ 
		Name: "Repair Unit",
		Level: "1d10",
		Type: []string{
			"Type: Wearable - Shoulder- or arm-mounted launcher,",
			"Usable: Handheld device",
		},
		Effect: `Device becomes a multiarmed sphere that
floats. It repairs one designated numenera
device (of a level equal to or less than its own)
that has been damaged but not destroyed. The
repair unit can even create spare parts, unless
the GM rules that the parts are too specialized
or rare (in which case, the unit repairs the
device entirely except for the specialized part).
Repair time is 1d100 + 20 minutes.`,
	},
	Cypher{ 
		Name: "Retaliation Nodule",
		Level: "1d6",
		Type: []string{
			"Usable: Crystal nodule affixed to armor",
		},
		Effect: `For the next 28 hours, anyone striking
the armor the nodule is attached to
triggers a small burst of electricity that
inflicts 1 point of damage (2 points if the
cypher is level 4 or higher, 3 points if the
cypher is level 6 or higher). No action or
roll is required by the armor’s wearer.`,
	},
	Cypher{ 
		Name: "Sheen",
		Level: "1d6",
		Type: []string{
			"Internal: Pill, ingestible liquid",
			"Usable: Injector",
		},
		Effect: `For one week, the user’s cells are coated with
a protective veneer that resists damage (+1 to
Armor, or +2 to Armor if the cypher is level 5 or
higher) and eases Might defense rolls by two
steps. However, healing is more difficult during
this time; all recovery rolls suffer a –1 penalty.`,
	},
	Cypher{ 
		Name: "Shock Nodule",
		Level: "1d6",
		Type: []string{
			"Usable: Crystal nodule affixed to a melee weapon",
		},
		Effect: `For the next 28 hours, each time the weapon
the nodule is attached to strikes a solid creature
or object, it generates a burst of electricity,
inflicting 1 additional point of damage (2 points
if the cypher is level 4 or higher, 3 points if
the cypher is level 6 or higher).`,
	},
	Cypher{ 
		Name: "Shocker",
		Level: "1d6 + 4",
		Type: []string{
			"Internal: Subdermal implant",
			"Type: Wearable - Ring, palm disk",
			"Usable: Short rod",
		},
		Effect: `Delivers a powerful burst of electricity
that shocks any creature touched, inflicting
damage equal to the cypher’s level.`,
	},
	Cypher{ 
		Name: "Skill Boost",
		Level: "1d6",
		Type: []string{
			"Internal: Pill, ingestible liquid",
			"Usable: Injector",
		},
		Effect: `Dramatically but temporarily alters the
user’s mind and body so that one specific
physical action they can perform is eased by
three steps. Once activated, this boost can be
used a number of times equal to the cypher’s
level, but only within a 28-hour period. The
boost takes effect each time the action is
performed, so a level 3 cypher boosts the
first three times the action is attempted. The
action can be one of a number of possibilities:
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
00 Escaping`,
	},
	Cypher{ 
		Name: "Sleep Inducer",
		Level: "1d6",
		Type: []string{
			"Internal: Pill, ingestible liquid",
			"Type: Wearable - Fingertip cusp, ring, glove",
			"Usable: Injector, gas sprayer",
		},
		Effect: `Touch or ingestion puts the victim to
sleep for ten minutes or until awoken by a
violent action or an extremely loud noise.`,
	},
	Cypher{ 
		Name: "Sonic Hole",
		Level: "1d6 + 2",
		Type: []string{
			"Type: Wearable - Wristband, ring, belt-mounted device",
			"Usable: Small handheld device",
		},
		Effect: `Draws all sound within long range into the
device for one round per cypher level. Within
the affected area, no sound can be heard.`,
	},
	Cypher{ 
		Name: "Sound Dampener",
		Level: "1d6 + 2",
		Type: []string{
			"Internal: Implant",
			"Type: Wearable - Wristband, ring, belt-mounted",
			"Usable: Small handheld device",
		},
		Effect: `Dampens all sound within immediate
range for one minute per cypher level,
providing an asset for all creatures in the
area to attempt stealthy actions.`,
	},
	Cypher{ 
		Name: "Spatial Warp",
		Level: "5",
		Type: []string{
			"Usable: Small metal ring",
		},
		Effect: `When affixed to another numenera
device that affects a single target at range,
that range is increased to 1 mile (1.5 km)
with no penalties. Space is temporarily
warped in terms of seeing and reaching the
target. If direct line of sight is important to
the device’s effect, it remains important.
Creating the spatial warp functions as one
use of the device.`,
	},
	Cypher{ 
		Name: "Speed Boost",
		Level: "1d6 + 2",
		Type: []string{
			"Internal: Pill, ingestible liquid",
			"Type: Wearable - Adhesive patch that activates when",
			"Usable: Injector",
		},
		Effect: `Substance adds 1 to Speed Edge for
one hour (or adds 2 if the cypher is level
5 or higher).`,
	},
	Cypher{ 
		Name: "Stim",
		Weight: 2,
		Level: "6",
		Type: []string{
			"Internal: Pill, ingestible liquid",
			"Usable: Injector",
		},
		Effect: `Eases the next action taken by three steps.`,
	},
	Cypher{ 
		Name: "Strength Boost",
		Level: "1d6 + 2",
		Type: []string{
			"Internal: Pill, ingestible liquid",
			"Type: Wearable - Adhesive patch that activates when",
			"Usable: Injector",
		},
		Effect: `Substance adds 1 to Might Edge for one
hour (or adds 2 if the cypher is level 5 or
higher).`,
	},
	Cypher{ 
		Name: "Subdual Field",
		Level: "1d6 + 3",
		Type: []string{
			"Usable: Complex device",
		},
		Effect: `Two rounds after being activated, the
device creates an invisible field that fills a
specified area (such as a cube of a certain
size) within long range of the device.
The field lasts for one minute. It affects
the minds of thinking beings within the
field, preventing them from taking hostile
actions. The effect lasts as long as they
remain in the field and for 1d6 rounds after,
although an Intellect defense roll is allowed
each round to act normally (both in the
field and after leaving it).`,
	},
	Cypher{ 
		Name: "Telepathy Implant",
		Level: "1d6 + 2",
		Type: []string{
			"Internal: Pill",
			"Type: Wearable - Disk that adheres to forehead,",
			"Usable: Injector",
		},
		Effect: `The user activates the device and
targets one creature within close range.
For one hour per cypher level, the device
enables two-way long-range mental
communication between the user and
the target. This lasts for one hour per
cypher level. Sometimes multiple cyphers
of this type are found together and allow
communication between all of them.`,
	},
	Cypher{ 
		Name: "Teleporter (bounder)",
		Level: "1d6 + 2",
		Type: []string{
			"Type: Wearable - Belt, wristband, ring, full bodysuit",
			"Usable: Complex device, handheld device",
		},
		Effect: `User teleports up to 100 × the cypher
level in feet (30 × cypher level in m) to a
location they can see. They arrive safely
with their possessions but cannot take
anything else with them.`,
	},
	Cypher{ 
		Name: "Teleporter (traveler)",
		Level: "1d6 + 4",
		Type: []string{
			"Type: Wearable - Belt, wristband, ring, full bodysuit",
			"Usable: Complex device, handheld device",
		},
		Effect: `User teleports up to 100 × the cypher
level in miles (160 x the cypher level in km)
to a location they have previously visited.
They arrive safely with their possessions
but cannot take anything else with them.`,
	},
	Cypher{ 
		Name: "Temporal Viewer",
		Level: "1d6 + 4",
		Type: []string{
			"Type: Wearable - Wristband",
			"Usable: Complex device, handheld device",
		},
		Effect: `Displays moving images and sound, up
to ten minutes per cypher level in length,
depicting events that occurred at the
current location up to one year prior. The
user specifies the time period shown by the
viewer.`,
	},
	Cypher{ 
		Name: "Time Dilation Nodule (defensive)",
		Level: "1d6",
		Type: []string{
			"Usable: Crystal nodule affixed to armor",
		},
		Effect: `For the next 28 hours, the wearer of the
armor moves in seemingly random, rapid
jumps, a few inches to one side or the
other, when attacked. This is an asset that
eases attacks by two steps (three steps if
the cypher is level 6 or higher).`,
	},
	Cypher{ 
		Name: "Time Dilation Nodule (offensive)",
		Level: "1d6",
		Type: []string{
			"Usable: Crystal nodule affixed to a melee",
		},
		Effect: `For the next 28 hours, the attacker
moves at almost instantaneous speeds
when they swing the weapon, easing their
attacks by two steps (three steps if the
cypher is level 6 or higher).`,
	},
	Cypher{ 
		Name: "Tracer",
		Level: "1d6",
		Type: []string{
			"Type: Wearable - Wristband",
			"Usable: Handheld device",
		},
		Effect: `Fires a microscopic tracer that clings to
any surface within short range (long range
if the cypher is level 4 or higher, very long
range if the cypher is level 6 or higher). For
the next 28 hours, the launcher shows the
distance and direction to the tracer, as long
as it is in the same dimension.`,
	},
	Cypher{ 
		Name: "Visage Changer",
		Level: "1d6",
		Type: []string{
			"Internal: Pill or injection that produces",
			"Usable: Tube of moldable paste",
		},
		Effect: `Changes the appearance of one humansized
creature, providing an asset to disguise
tasks (easing them by two steps if the cypher
is level 5 or higher). The change takes ten
minutes to apply and lasts for 28 hours.`,
	},
	Cypher{ 
		Name: "Visual Displacement Device",
		Level: "1d6",
		Type: []string{
			"Type: Wearable - Belt or bracelet",
			"Usable: Handheld device",
		},
		Effect: `Projects holographic images of the
wearer to confuse attackers. The images
appear around the wearer. This gives the
wearer an asset to Speed defense actions
for ten minutes per cypher level.`,
	},
	Cypher{ 
		Name: "Vocal Translator",
		Level: "1d6",
		Type: []string{
			"Internal: Pill, ingestible liquid",
			"Type: Wearable - Headband",
			"Usable: Disk that must be held to forehead",
		},
		Effect: `Translates everything said by the
user into a language that anyone can
understand for 28 hours per cypher level.`,
	},
	Cypher{ 
		Name: "Warmth Projector",
		Level: "1d6",
		Type: []string{
			"Internal: Subdermal implant",
			"Type: Wearable - Bodysuit, belt",
			"Usable: Injector",
		},
		Effect: `Keeps the user warm and comfortable
in the harshest cold temperatures for
28 hours. During this time, the user
has Armor equal to the cypher level that
protects against cold damage.`,
	},
	Cypher{ 
		Name: "Water Breather",
		Level: "1d6",
		Type: []string{
			"Internal: Pill, ingestible liquid",
			"Usable: Injector",
		},
		Effect: `Allows an air breather to extract oxygen
from water for five hours per cypher level
so they can breathe underwater.`,
	},
	Cypher{ 
		Name: "X-Ray Viewer",
		Level: "1d6 + 4",
		Type: []string{
			"Usable: Glass panel",
		},
		Effect: `When held up against a solid surface,
this panel allows the user to see through
up to 2 feet (60 cm) of material. The panel
works only if the cypher’s level is higher
than the material’s level. The effect lasts for
one minute per cypher level.`,
	},
	Cypher{ 
		Name: "Ablative Thrower",
		Level: "1d6 + 3",
		Type: []string{
			"Usable: Bulky handheld device attached to",
		},
		Effect: `When activated, creatures within short
range are coated with a fine mist of metallic
nano particles that coats their clothing (but
not flesh), hardening non-flexing portions
for about an hour and granting +1 Armor. If
the use is coordinated with attackers who
make up a community’s or horde’s combat
force, the force is treated as a defensive
horde during that community action.`,
	},
	Cypher{ 
		Name: "Acidic Resonator",
		Level: "1d6 + 2",
		Type: []string{
			"Usable: Bulky handheld device",
		},
		Effect: `When activated, any exposed metallic
weapons within short range bead with
acid for about an hour, granting affected
weapons 1 additional point of damage
when used in combat. If the use is
coordinated with attackers who make up a
community’s or horde’s combat force, the
community or horde inflicts +1 damage
during the next community-scale combat
interaction.`,
	},
	Cypher{ 
		Name: "Community Beacon",
		Level: "1d6 + 2",
		Type: []string{
			"Usable: Handheld device",
		},
		Effect: `When activated, the cypher splits
into two pieces. One piece is affixed to a
numenera structure, and the other is kept
by a character. The character can then use
the connection between the two pieces to
trigger a spatial warp any time within the
next 28 hours. The spatial warp returns
the user and all targets within immediate
range to a location next to the numenera
structure, no matter how far apart the
character and the structure were.`,
	},
	Cypher{ 
		Name: "Crafter’s Eyes",
		Level: "1d6 + 2",
		Type: []string{
			"Wearable: Thick lenses on frames worn over",
		},
		Effect: `Informative images formed on the inner
lenses allow the user to automatically
succeed on any one crafting subtask whose
level is equal to or less than the cypher’s
level.`,
	},
	Cypher{ 
		Name: "Deception Filter",
		Level: "1d6",
		Type: []string{
			"Internal: Pill, ingestible liquid",
			"Wearable: Badge",
		},
		Effect: `User cannot knowingly tell a lie for one
hour.`,
	},
	Cypher{ 
		Name: "Destiny Nodule",
		Level: "1d6",
		Type: []string{
			"Usable: Crystal nodule affixed to tool",
		},
		Effect: `For the next 28 hours, each time the
tool the nodule is attached to is used
to repair, craft, or modify an object or
structure, the tool provides the user an
additional asset to the task.`,
	},
	Cypher{ 
		Name: "Ablative Thrower",
		Level: "1d6 + 3",
		Type: []string{
			"Usable: Bulky handheld device attached to",
		},
		Effect: `When activated, creatures within short
range are coated with a fine mist of metallic
nano particles that coats their clothing (but
not flesh), hardening non-flexing portions
for about an hour and granting +1 Armor. If
the use is coordinated with attackers who
make up a community’s or horde’s combat
force, the force is treated as a defensive
horde during that community action.`,
	},
	Cypher{ 
		Name: "Detonation",
	},
	Cypher{ 
		Name: "(horde Suppressor)",
		Level: "1d6 + 1",
		Type: []string{
			"Wearable: Wristband projector (long range)",
			"Usable: Handheld projector (long range)",
		},
		Effect: `Explodes to release a burst of neurotoxic
mist affecting all creatures within long
range, hindering all attacks for about an
hour. If effectively targeted during a conflict
against a ranked horde or community,
the affected horde or community inflicts 1
point of damage less than normal during
that community action.`,
	},
	Cypher{ 
		Name: "Effort Enhancement",
		Level: "1d6 + 2",
		Type: []string{
			"Internal: Pill, ingestible liquid",
			"Wearable: Adhesive patch that activates when",
			"Usable: Injector",
		},
		Effect: `For the next 28 hours, the user can
access a thin connection to the datasphere
to gain a very specific effect. When applying
Effort to a task related to a predetermined
skill, the user can apply one free level of
Effort. The skills encoded are generally
restricted to the following.
01–10 Melee attacks
11–20 Ranged attacks
21–40 Understanding numenera
41–50 Salvaging numenera
51–60 Crafting numenera
61–70 Persuasion
71–75 Healing
76–80 Speed defense
81–85 Intellect defense
86–90 Swimming
91–95 Riding
96–00 Sneaking`,
	},
	Cypher{ 
		Name: "Emotion Smoother",
		Level: "1d6",
		Type: []string{
			"Internal: Pill, ingestible liquid",
			"Wearable: Adhesive patch that activates when",
			"Usable: Injector",
		},
		Effect: `For the next hour, the character loses
access to the part of their mind that
registers and produces emotion. This
allows the user to experience situations
they might not otherwise be able to handle,
to fairly weigh options without emotional
shortcuts, or to tell completely convincing
lies because they are no longer generating
micro-expressions or other telling responses
due to emotional inconsistency.`,
	},
	Cypher{ 
		Name: "Farspeaker",
		Level: "1d6",
		Type: []string{
			"Usable: Handheld device",
		},
		Effect: `For the next ten minutes, the user can be
heard at any distant location that they can
see, as long as nothing physically blocks
the intervening space. They can also hear
sounds made in return. The area in which
the user can speak and hear is up to a
short distance across in the target location.`,
	},
	Cypher{ 
		Name: "Gravity Changer",
		Level: "1d6 + 2",
		Type: []string{
			"Useable: Metal plates",
		},
		Effect: `Once placed, the plates change the
direction of gravity within an area up to
a short distance in diameter. The gravitic
warp could be continuous across the
affected area, or it could bend and curve,
changing directions within the area up to
a number of times equal to the level of the
cypher. Once gravity is warped, it cannot
be changed and persists for 28 hours.`,
	},
	Cypher{ 
		Name: "Harmony Lamp",
		Level: "1d6",
		Type: []string{
			"Useable: Crystal orb, nodule, disc",
		},
		Effect: `For an hour after activation, the cypher
emits a combination of colored light, sound,
odor, and healing nanobots that permeate an
area up to a short distance across. All positive
interaction tasks made in the area gain an asset.
Creatures who’ve made up their minds about a
topic may be open to revisiting the issue. PCs
add +1 to recovery rolls made in the area.`,
	},
	Cypher{ 
		Name: "Health Resonator",
		Level: "1d6 + 2",
		Type: []string{
			"Usable: Bulky handheld device",
		},
		Effect: `When activated, creatures within short
range each gain 2 points of health (or 2
points to Might for PCs) for about an hour.
If the use is coordinated with attackers who
make up a community’s or horde’s combat
force, the force has +2 health during that
community action.`,
	},
	Cypher{ 
		Name: "Hiding Alarm Nodule",
		Level: "7",
		Type: []string{
			"Usable: Crystal nodule affixed to installation",
		},
		Effect: `For the next 28 hours, each time the
installation to which the nodule is attached
is struck hard enough to inflict damage (but
not more than once per round), the structure
or installation goes out of phase for one hour.
The installation must fit in a cube up to a short
distance wide, tall, and deep. All contents of
the out-of-phase structure also go out of phase,
unless they somehow leave the structure.`,
	},
	Cypher{ 
		Name: "Immobilizer",
		Level: "1d6 + 2",
		Type: []string{
			"Usable: Handheld device",
		},
		Effect: `The user directs one end of the cypher
at a target within short range, which is then
subject to an attack by a projected mass of
rapidly expanding and hardening foam. The
target is held immobile in the hardened
mass of foam for one hour, after which the
mass turns to powder.`,
	},
	Cypher{ 
		Name: "Infrastructure Drill",
		Level: "1d6 + 1",
		Type: []string{
			"Usable: Large sphere with a thick screw protrusion",
		},
		Effect: `The user attaches this cypher to a wall
or other structure. On the following round,
it drills into the object and keeps going
beneath the surface for the next hour,
creating a series of tunnels that weaken the
integrity of the structure and connecting
structures. This decreases the level of all
commonplace structures connected to
the initial structure within long range and
inflicts 3 points of damage to a ranked
community’s infrastructure.`,
	},
	Cypher{ 
		Name: "Installation Enhancer",
		Level: "1d6 + 3",
		Type: []string{
			"Wearable: Bulky device and several metal plates",
		},
		Effect: `When securely attached to an
installation, it increases the level of the
entire installation up to the level of the
cypher for about one hour (no effect on
installations of a level equal to or higher
than the cypher).`,
	},
	Cypher{ 
		Name: "Installation Mover",
		Level: "1d6 + 2",
		Type: []string{
			"Wearable: Bulky device and several metal plates",
		},
		Effect: `When securely attached to an installation,
it allows the entire installation to be moved
to another location within very long range
(which is normally a difficult task). The
transfer requires that the user be able to push
the installation to the new location, but for
the duration of about an hour, the installation
seems almost weightless. During the transfer,
the installation does not function.`,
	},
	Cypher{ 
		Name: "Instant Item",
		Level: "1d6",
		Type: []string{
			"Wearable: Bulky amulet",
		},
		Effect: `This cypher has a two-stage activation.
The first stage occurs when the wearer keys
an object that they can hold in one hand to
the cypher. The object falls into a subspace
pocket and remains indefinitely or until a
user activates the cypher a second (and final)
time, immediately retrieving the stored item.`,
	},
	Cypher{ 
		Name: "Instant Wall",
		Level: "1d6 + 1",
		Type: []string{
			"Usable: Handheld device",
		},
		Effect: `With the addition of water and air, the
small device expands into an immobile
plane up to 10 feet (3 m) tall, 30 feet (9
m) long, and 2 feet (60 cm) thick. It is
made from a form of shapestone and is
permanent and immobile once created.`,
	},
	Cypher{ 
		Name: "Instant Workshop",
		Level: "1d6 + 2",
		Type: []string{
			"Usable: Handheld device",
		},
		Effect: `This small device expands into a
workbench that fixes itself in place. The
workbench provides an adjustable-height
work surface with light and a variety
of attached tools suitable for working
with commonplace objects, numenera
objects, and structural components.
The workshop enables crafting tasks
by providing the appropriate tools and
workspace, though it does not provide
shelter from the elements. The workbench
and all its components are made from a
form of shapemetal, and the workbench is
permanent and immobile once created.`,
	},
	Cypher{ 
		Name: "Iotum Stabilizer",
		Level: "1d6 + 1",
		Type: []string{
			"Usable: Bulky handheld device",
		},
		Effect: `When used as part of a salvaging task,
the field emitted by the cypher stabilizes
all iotum within short range so that an
additional 1d6 units of iotum are recovered.
The additional iotum must be of the same
kind that has already been discovered, and
the iotum must be of a level equal to or less
than the level of this cypher.`,
	},
	Cypher{ 
		Name: "Iotum Upgrader",
		Level: "1d6 + 2",
		Type: []string{
			"Usable: Bulky handheld device",
		},
		Effect: `When attached to 1 unit of iotum, that
iotum can then be used as if it were a unit
of iotum 1 level higher in a crafting task.
For example, an upgraded unit of io (level 1)
could be used as if it were a unit of
responsive synth (level 2) for the purpose
of one crafting task, or a unit of quantium
(level 5) could be used as if it were a unit of
protomatter (level 6).`,
	},
	Cypher{ 
		Name: "Iron Wind Resistance",
		Level: "7",
		Type: []string{
			"Internal: Pill, ingestible liquid",
			"Usable: Injector",
		},
		Effect: `Grants the user two assets to defense
tasks made to resist the effects of an
instance of the iron wind (and attempts
to resist any ongoing effects instigated by
previous iron wind contact, if any) for the
next 28 hours.`,
	},
	Cypher{ 
		Name: "Mechanical Wright",
		Level: "1d6 + 3",
		Type: []string{
			"Usable: Complex device",
		},
		Effect: `Crafts one numenera object or
structure whose level can be no higher
than the mechanical wright’s level
minus 3 (minimum level 1), as long as a
numenera plan is provided to work from.
Crafting occurs over the course of the
next 28 hours, regardless of the target
device’s level. Components, including
parts and iotum, are contributed from
the mechanical wright, which becomes
completely integrated into the new
numenera device over the course of the
crafting process.`,
	},
	Cypher{ 
		Name: "Object Replicator",
		Level: "1d6 + 1",
		Type: []string{
			"Usable: Crystal nodule affixed to object that",
		},
		Effect: `Replicates three additional copies of
most objects whose level is equal to or less
than this cypher’s level. A unit of iotum can
be replicated, but not functioning cyphers,
artifacts, or similar devices.`,
	},
	Cypher{ 
		Name: "Organ Factory",
		Level: "1d6",
		Type: []string{
			"Internal: Pill, ingestible liquid",
			"Usable: Injector",
		},
		Effect: `Over the course of 28 hours, the
user grows an external, fully viable living
organ or limb. This is a somewhat painful
process that dazes the user for the last
half of the process. The completely grown
organ or limb can be transferred to another
creature simply by placing it on their body
and waiting a few rounds. If the recipient
creature has sustained damage to the
same kind of organ or limb, the newly
grown one replaces it over the course of
about a minute (during which time the
recipient creature is stunned). Otherwise,
the duplicate organ or limb withers and
dies.
Duplicate organs must be used within one
week. All force-grown organs and limbs are
sealed within a partially living, protective,
translucent caul. The cypher can produce
one of the following, chosen by the user at
the time of the cypher’s use:
01–04 Heart
05–08 Lung 
09–12 Brain 
13–16 Blood 
17–20 Ribs 
21–25 Eye 
26–30 Ear 
31–35 Hand
36–45 Foot
46–55 Full arm
56–65 Full leg
66–75 Nose
76–85 Mouth
86–95 Stomach
96–00 Intestines`,
	},
	Cypher{ 
		Name: "Phasing Gloves",
		Level: "1d6 + 1",
		Type: []string{
			"Wearable: Gloves",
		},
		Effect: `These gloves can selectively phase
through solid objects whose level is no
higher than the cypher level in a fashion
that allows the wearer to see what they are
touching, providing an asset for anyone
attempting to repair a disabled object or
structure. Someone wearing the gloves
could also search behind walls and under
floors using the gloves. The gloves retain
their phasing ability for up to one minute per
cypher level. If used as part of a salvaging
task, the gloves grant two assets.`,
	},
	Cypher{ 
		Name: "Psychic Defense Nodule",
		Level: "1d6 + 2",
		Type: []string{
			"Usable: Crystal nodule affixed to side of head",
		},
		Effect: `For the next 28 hours, each time the
wearer of the nodule is affected by an
attack that attempts to sway their actions
or beliefs or that inflicts Intellect damage,
they instead go into stasis for one round
and remain unaffected. While in stasis,
they also lose their next turn. Attacks
against a wearer in stasis are made as if
the wearer were protected by a force field
whose level is equal to the cypher’s level.`,
	},
	Cypher{ 
		Name: "Rage Thrower",
		Level: "1d6 + 2",
		Type: []string{
			"Usable: Bulky handheld device attached to",
		},
		Effect: `When activated, creatures within
short range are coated with a fine mist of
insinuating nano particles that enter their
blood and stimulate their behavior for
about an hour, making them more violent
but less able to feel pain. This grants them
+1 Armor but an inability to use Effort
from their Intellect Pool. If the use is
coordinated with attackers who make up
a community’s or horde’s combat force,
the force is treated as a marauding horde
during that community action.`,
	},
	Cypher{ 
		Name: "Retriever",
		Level: "1d6 + 3",
		Type: []string{
			"Wearable: Glove of synth and small nodule",
			"Usable: Small handheld device and small",
		},
		Effect: `These cyphers always come in at
least two parts: a nodule and some kind
of retrieving device. The nodule can be
attached to an item your size or smaller.
When the cypher is activated, that item is
teleported to the device.`,
	},
	Cypher{ 
		Name: "Salvage Graft",
		Level: "1d6 + 1",
		Type: []string{
			"Usable: Small metallic disc",
		},
		Effect: `When applied to a creature’s head, the
disk immediately unleashes microfilaments
that enter the brain. Within five minutes,
the creature is trained in tasks related to
salvaging numenera. If the creature is
already trained or specialized in salvaging
numenera, this graft has no effect. Once
the graft attaches, the effect is permanent,
and this device no longer counts against
the number of cyphers that a PC can bear,
but the disc remains.`,
	},
	Cypher{ 
		Name: "Shapemetal",
		Level: "6",
		Type: []string{
			"Usable: Canister containing a silvery,",
		},
		Effect: `This metallic clay can be shaped and
greatly expanded to create individual tools
or objects such as a hammer, a ladder, or
a basic structure such as a wall, bench,
floor, staircase, and so on, as long as the
total volume created could fit in a 10-foot
(3 m) cube. Once formed, the structure or
objects are permanent.`,
	},
	Cypher{ 
		Name: "Slash-Retardant Spray",
		Level: "1d6 + 2",
		Type: []string{
			"Usable: Bulky handheld device attached to",
		},
		Effect: `For the next 28 hours, walls and
structures sprayed with these repairing
nanites are treated as if 1 level higher. If
an hour is spent spraying down the outer
walls and defense structures of a ranked
community, that community gains +1
Armor during any conflict that happens
over the next 28 hours.`,
	},
	Cypher{ 
		Name: "Stealth Thrower",
		Level: "1d6 + 2",
		Type: []string{
			"Usable: Bulky handheld device attached to",
		},
		Effect: `When activated, creatures within
short range are coated with a fine mist
of dull nano particles that render them
more difficult to pick out from their
surroundings for about an hour, granting
them an asset to stealth tasks. If the use is
coordinated with attackers who make up a
community’s or horde’s combat force, the
force is treated as a stealthy horde during
that community action.`,
	},
	Cypher{ 
		Name: "Summoning Alarm Nodule",
		Level: "6",
		Type: []string{
			"Usable: Crystal nodule affixed to installation",
		},
		Effect: `For the next 28 hours, each time the
structure or installation to which the nodule
is attached is struck hard enough to inflict
damage (but not more than once per round),
whoever attached the nodule is teleported
from any location within very long range to a
location standing next to the installation.`,
	},
	Cypher{ 
		Name: "Sustenance Implant",
		Level: "1d6",
		Type: []string{
			"Internal: Pill, ingestible liquid",
			"Usable: Injector",
		},
		Effect: `For the next number of days equal to
the cypher level, the wearer feels no ill
effects from not eating or from overeating.
They also gain an asset to any defense
tasks to withstand poison.`,
	},
	Cypher{ 
		Name: "Trail Blazer",
		Level: "1d6",
		Type: []string{
			"Wearable: Single boot with attached device",
		},
		Effect: `For the next 28 hours, each time the boot
touches down on solid ground, it deposits
a patch of nanites that mark the area with
a symbol. At any time before the duration
expires, the wearer (and up to one other
person that can be carried along) can use
an action to instantly teleport back to any
of these symbols. The teleportation effect
brings the cypher’s duration to an end.`,
	},
	Cypher{ 
		Name: "Transdimensional Sense",
		Level: "1d6",
		Type: []string{
			"Internal: Pill, ingestible liquid",
			"Usable: Injector",
		},
		Effect: `The user can see out-of-phase, invisible,
and transdimensional creatures, objects, and
sources of energy within long range for 28
hours. If ingested by a creature who attempts
to salvage a particular kind of iotum during
a salvage task, the effects of this cypher
grant the user one free level of Effort if they
first use a level of Effort on that task.`,
	},
	Cypher{ 
		Name: "Travel Bubble",
		Level: "1d6 + 4",
		Type: []string{
			"Useable: Handheld device",
		},
		Effect: `The user and any additional targets that
can fit into a space an immediate distance
in diameter are encapsulated in a spherical
force field for up to ten hours or until the
user collapses it. The sphere hovers at all
times and moves as the user wishes vertically
or horizontally up to a short distance each
round. If the sphere is destroyed or collapsed,
whatever it contains falls to the ground.`,
	},
}
