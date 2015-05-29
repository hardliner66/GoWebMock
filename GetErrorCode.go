package main

import (
	"math/rand"
	"time"
)

type HttpErrorCode struct {
	Code        int
	Description string
}

func GetErrorCode() HttpErrorCode {
	InitializeErrorCodes()
	return errorCodes[randInt(0, len(errorCodes)-1)]
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

var errorCodes []HttpErrorCode
var initialized bool = false

func InitializeErrorCodes() {
	if !initialized {
		rand.Seed(time.Now().UTC().UnixNano())
		errorCodes = []HttpErrorCode{
			//* 70X - Inexcusable
			HttpErrorCode{701, "Meh"},
			HttpErrorCode{702, "Emacs"},
			HttpErrorCode{703, "Explosion"},
			HttpErrorCode{704, "Goto Fail"},
			HttpErrorCode{705, "I wrote the code and missed the necessary validation by an oversight (see 795)"},
			//* 71X - Novelty Implementations
			HttpErrorCode{710, "PHP"},
			HttpErrorCode{711, "Convenience Store"},
			HttpErrorCode{712, "NoSQL"},
			HttpErrorCode{719, "I am not a teapot"},
			//* 72X - Edge Cases
			HttpErrorCode{720, "Unpossible"},
			HttpErrorCode{721, "Known Unknowns"},
			HttpErrorCode{722, "Unknown Unknowns"},
			HttpErrorCode{723, "Tricky"},
			HttpErrorCode{724, "This line should be unreachable"},
			HttpErrorCode{725, "It works on my machine"},
			HttpErrorCode{726, "It's a feature}, not a bug"},
			HttpErrorCode{727, "32 bits is plenty"},
			//* 73X - Fucking
			HttpErrorCode{730, "Fucking Bower"},
			HttpErrorCode{731, "Fucking Rubygems"},
			HttpErrorCode{732, "Fucking Unic💩de"},
			HttpErrorCode{733, "Fucking Deadlocks"},
			HttpErrorCode{734, "Fucking Deferreds"},
			HttpErrorCode{735, "Fucking IE"},
			HttpErrorCode{736, "Fucking Race Conditions"},
			HttpErrorCode{737, "FuckThreadsing"},
			HttpErrorCode{738, "Fucking Bundler"},
			HttpErrorCode{739, "Fucking Windows"},
			//* 74X - Meme Driven
			HttpErrorCode{740, "Computer says no"},
			HttpErrorCode{741, "Compiling"},
			HttpErrorCode{742, "A kitten dies"},
			HttpErrorCode{743, "I thought I knew regular expressions"},
			HttpErrorCode{744, "Y U NO write integration tests?"},
			HttpErrorCode{745, "I don't always test my code}, but when I do I do it in production"},
			HttpErrorCode{746, "Missed Ballmer Peak"},
			HttpErrorCode{747, "Motherfucking Snakes on the Motherfucking Plane"},
			HttpErrorCode{748, "Confounded by Ponies"},
			HttpErrorCode{749, "Reserved for Chuck Norris"},
			//* 75X - Syntax Errors
			HttpErrorCode{750, "Didn't bother to compile it"},
			HttpErrorCode{753, "Syntax Error"},
			HttpErrorCode{754, "Too many semi-colons"},
			HttpErrorCode{755, "Not enough semi-colons"},
			HttpErrorCode{756, "Insufficiently polite"},
			HttpErrorCode{757, "Excessively polite"},
			HttpErrorCode{759, "Unexpected T_PAAMAYIM_NEKUDOTAYIM"},
			//* 76X - Substance-Affected Developer
			HttpErrorCode{761, "Hungover"},
			HttpErrorCode{762, "Stoned"},
			HttpErrorCode{763, "Under-Caffeinated"},
			HttpErrorCode{764, "Over-Caffeinated"},
			HttpErrorCode{765, "Railscamp"},
			HttpErrorCode{766, "Sober"},
			HttpErrorCode{767, "Drunk"},
			HttpErrorCode{768, "Accidentally Took Sleeping Pills Instead Of Migraine Pills During Crunch Week"},
			HttpErrorCode{769, "Questionable Maturity Level"},
			//* 77X - Predictable Problems
			HttpErrorCode{771, "Cached for too long"},
			HttpErrorCode{772, "Not cached long enough"},
			HttpErrorCode{773, "Not cached at all"},
			HttpErrorCode{774, "Why was this cached?"},
			HttpErrorCode{776, "Error on the Exception"},
			HttpErrorCode{777, "Coincidence"},
			HttpErrorCode{778, "Off By One Error"},
			HttpErrorCode{779, "Off By Too Many To Count Error"},
			//* 78X - Somebody Else's Problem
			HttpErrorCode{780, "Project owner not responding"},
			HttpErrorCode{781, "Operations"},
			HttpErrorCode{782, "QA"},
			HttpErrorCode{783, "It was a customer request}, honestly"},
			HttpErrorCode{784, "Management}, obviously"},
			HttpErrorCode{785, "TPS Cover Sheet not attached"},
			HttpErrorCode{786, "Try it now"},
			HttpErrorCode{787, "Further Funding Required"},
			HttpErrorCode{788, "Designer's final designs weren't"},
			//* 79X - Internet crashed
			HttpErrorCode{791, "The Internet shut down due to copyright restrictions."},
			HttpErrorCode{792, "Climate change driven catastrophic weather event"},
			HttpErrorCode{793, "Zombie Apocalypse"},
			HttpErrorCode{794, "Someone let PG near a REPL"},
			HttpErrorCode{795, "#heartbleed (see 705)"},
			HttpErrorCode{797, "This is the last page of the Internet. Go back"},
			HttpErrorCode{799, "End of the world"},
		}
		initialized = true
	}
}
