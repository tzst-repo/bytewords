package bytewords

import (
	"strings"
)

var (
	// bytewordList is the set of bytewords to use.
	bytewordList []string

	// bytewordMap is a reverse lookup map for bytewordList.
	bytewordMap map[string]int

	// shortBytewordMap is a reverse lookup map for shortBytewordS.
	shortBytewordMap map[string]int
)

func init() {
	SetBytewordList(Bytewords)
}

func SetBytewordList(list []string) {
	bytewordList = list
	bytewordMap = map[string]int{}
	shortBytewordMap = map[string]int{}

	for i, v := range bytewordList {
		bytewordMap[v] = i
	}

	for i, v := range bytewordList {
		shortBytewordMap[GetShortByteWord(v)] = i
	}

}

func EncodeShortBytewords(bytes []byte) string {
	var shortbytewords string
	for _, hex := range bytes {
		byteword := bytewordList[int(hex)]
		shortbytewords = shortbytewords + GetShortByteWord(byteword)
	}
	return shortbytewords
}

func GetShortByteWord(byteword string) string {
	return byteword[0:1] + byteword[3:4]
}

func DecodeShortBytewords(shortBytewords string) []byte {
	// slice to hold bytes in
	chunk := ConvertShortByteWordStringtoSlice(shortBytewords)
	var bytes []byte
	bytes = make([]byte, len(chunk))

	for i, short := range chunk {
		output := shortBytewordMap[short]
		bytes[i] = byte(output)
	}
	return bytes
}

func ConvertShortByteWordStringtoSlice(shortByteWords string) []string {
	var shortByteWordSlice []string
	shortByteWord := make([]rune, 2)
	len := 0
	for _, r := range shortByteWords {
		shortByteWord[len] = r
		len++
		if len == 2 {
			shortByteWordSlice = append(shortByteWordSlice, string(shortByteWord))
			len = 0
		}
	}
	if len > 0 {
		shortByteWordSlice = append(shortByteWordSlice, string(shortByteWord[:len]))
	}
	return shortByteWordSlice
}

var Bytewords = strings.Split(strings.TrimSpace(bytewords), "\n")
var bytewords = `
able
acid
also
apex
aqua
arch
atom
aunt
away
axis
back
bald
barn
belt
beta
bias
blue
body
brag
brew
bulb
buzz
calm
cash
cats
chef
city
claw
code
cola
cook
cost
crux
curl
cusp
cyan
dark
data
days
deli
dice
diet
door
down
draw
drop
drum
dull
duty
each
easy
echo
edge
epic
even
exam
exit
eyes
fact
fair
fern
figs
film
fish
fizz
flap
flew
flux
foxy
free
frog
fuel
fund
gala
game
gear
gems
gift
girl
glow
good
gray
grim
guru
gush
gyro
half
hang
hard
hawk
heat
help
high
hill
holy
hope
horn
huts
iced
idea
idle
inch
inky
into
iris
iron
item
jade
jazz
join
jolt
jowl
judo
jugs
jump
junk
jury
keep
keno
kept
keys
kick
kiln
king
kite
kiwi
knob
lamb
lava
lazy
leaf
legs
liar
limp
lion
list
logo
loud
love
luau
luck
lung
main
many
math
maze
memo
menu
meow
mild
mint
miss
monk
nail
navy
need
news
next
noon
note
numb
obey
oboe
omit
onyx
open
oval
owls
paid
part
peck
play
plus
poem
pool
pose
puff
puma
purr
quad
quiz
race
ramp
real
redo
rich
road
rock
roof
ruby
ruin
runs
rust
safe
saga
scar
sets
silk
skew
slot
soap
solo
song
stub
surf
swan
taco
task
taxi
tent
tied
time
tiny
toil
tomb
toys
trip
tuna
twin
ugly
undo
unit
urge
user
vast
very
veto
vial
vibe
view
visa
void
vows
wall
wand
warm
wasp
wave
waxy
webs
what
when
whiz
wolf
work
yank
yawn
yell
yoga
yurt
zaps
zero
zest
zinc
zone
zoom
`
