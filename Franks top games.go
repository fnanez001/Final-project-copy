// Frank Nanez
// 4-25-20
// program to display my top 25 games played on steam and give information for user to learn more about the game.

//load main package and import fmt, math/rand, time, and strconv packages into program.
package main

import (
  "fmt"
  "math/rand"
  "time"
  //"strconv"
)

// create a struct for game info objects to include number, hours played, name, genre, and synopsis
type ginfo struct {
  Number  int
  Hrs float64
  Name, Genre, Synopsis string

}

// create a function that can pick a random number and can convert to string to use in if statements later
func randPick(answer int) (int) {
 rand.Seed(time.Now().UnixNano())
 n := (rand.Intn(24))
 //number := strconv.Itoa(n)
 return n
}

// create your main func inlcude all objects and classes for games and info needed.
func main (){
 var game[25]ginfo
  game[0] = ginfo{
    Number: 1,
    Name: "Banished",
    Hrs: 349,
    Genre: "Simulation, Stragey, City Builder",
    Synopsis: "You and a group of exiled travelers who decide to restart their lives in a new land. They have only the clothes on their backs and a cart filled with supplies from their homeland. The townspeople of Banished are your primary resource. You must build, survive, and if you’re good enough, thrive.",
  }
  game[1] = ginfo{
    Number: 2,
    Name:	"Dota 2",
    Hrs: 142,
    Genre: "MOBA, strategy",
    Synopsis: "Primary game mode is player vs player (PvP) where two teams of five battle against one another with waves of computer minions to help you push a “lane” guarded by other players, their minions, and towers.",
  }
  game[2] = ginfo{
    Number: 3,
    Name: "Mini Ninja’s",
    Hrs: 94,
    Genre: "action, adventure, indie",
    Synopsis: "Mini Ninjas is a game that combines furious action with stealth and exploration for an experience that appeals to a wide audience across age groups and preferences. Set with a ninja oriental theme. Fight against the darkness that is attacking your home and the ones you know and love.",
  }
  game[3] = ginfo{
    Number: 4,
    Name: "The Elder Scrolls V: Skyrim",
    Hrs: 79,
    Genre: "Open world, RPG, first person",
    Synopsis: "Dragons were thought to be extinct. You, a prisoner, survive and attack by a dragon and find out that you are a “dragonborn”. One who can speak the dragon language and use their shouts, a special ability/magic. You can create a character with tons of customization options and explore and play the game anyway you desire with the vast open world with plenty of choice and quests to take.",
  }
  game[4] = ginfo{
    Number: 5,
    Name: "Sid Meier's Civilization V",
    Hrs: 62,
    Genre: "Turn-based strategy, 4x",
    Synopsis: "Play one of history’s greatest civilizations. Employ diplomatic, research, and military strategies to be the best civilization and win, choosing from many maps and map sizes for the set of your theatre.",
  }
  game[5] = ginfo{
    Number: 6,
    Name: "Borderlands 2",
    Hrs: 55,
    Genre: "open world, RPG, first person",
    Synopsis: "Play as one of four vault hunters facing off against a massive world of creatures, psychos and the evil mastermind, Handsome Jack. Make new friends, arm them with a bazillion weapons to help you on your adventure.",
  }
  game[6] = ginfo{
    Number: 7,
    Name: "Company of Heroes-Legacy edition",
    Hrs: 51,
    Genre: "real-time strategy",
    Synopsis: "Delivering a visceral WWII gaming experience, Company of Heroes redefines real time strategy gaming by bringing the sacrifice of heroic soldiers, war-ravaged environments, and dynamic battlefields to life. Beginning with the D-Day Invasion of Normandy, players lead squads of Allied soldiers into battle against the German war machine through some of the most pivotal battles of WWII. Through a rich single player campaign, players experience the cinematic intensity and bravery of ordinary soldiers thrust into extraordinary events.",
  }
  game[7] = ginfo{
    Number: 8,
    Name: "PAYDAY 2",
    Hrs: 47,
    Genre: "First person shooter, strategy/heist",
    Synopsis: "Team up with up to four friends or other online players to take on heists and earn money, buy weapons, and customize your gear to help you be successful.",
  }
  game[8] = ginfo{
    Number: 9,
    Name: "The Forest",
    Hrs: 47,
    Genre: "Survival, horror, PVE",
    Synopsis: "Your plane crash lands on this beautiful island with a lush environment. You awake after the crash landing, only remembering your son calling out to you as he is dragged away, having blurry vision you don’t remember many details. You need to build some shelter and a base camp for your exploration. Soon after establishing some semblance of a shelter you see something off in the distance scoping out your new home. It’s a savage tribal inhabitant that you soon find out is a cannibal. You need to hurry and find your son and find a way off the island.",
  }
  game[9] = ginfo{
    Number: 10,
    Name: "Endless Legend",
    Hrs: 45,
    Genre: "Turn-based strategy, 4X",
    Synopsis: ": Play as one of many civilizations, or create one of your own. Employ diplomatic, research, and military strategies to be the best civilization and win. Select diverse maps and setting to help make every skirmish unique and challenging.",
  }
  game[10] = ginfo{
    Number: 11,
    Name: "Unturned",
    Hrs: 34,
    Genre: "survival, PVE, PVP",
    Synopsis: "Create or join a world in which you survive the zombie apocalypse.",
  }
  game[11] = ginfo{
    Number: 12,
    Name: "Tales of Zestiria",
    Hrs: 33,
    Genre: "action, RPG, anime style",
    Synopsis: "In a world torn by war and political skirmishes, accept the burden of the Shepherd and fight human darkness to protect your world from Malevolence and reunite humans and Seraphim.",
  }
  game[12] = ginfo{
    Number: 13,
    Name: "South Park™: The Stick of Truth™",
    Hrs: 33,
    Genre: "turn-based strategy, RPG",
    Synopsis: "From the perilous battlefields of the fourth-grade playground, a young hero will rise, destined to be South Park’s savior. From the creators of South Park, Trey Parker and Matt Stone, comes an epic quest to become… cool. Introducing South Park™: The Stick of Truth™.",
  }
  game[13] = ginfo{
    Number: 14,
    Name: "Endless Space",
    Hrs: 28,
    Genre: "turn-based strategy, 4X",
    Synopsis: "Covering the space colonization age in the Endless universe, where you can control every aspect of your civilization as you strive for galactic domination.",
  }
  game[14] = ginfo{
    Number: 15,
    Name: "Kerbal Space Program",
    Hrs: 27,
    Genre: "simulation, sandbox",
    Synopsis: ": In Kerbal Space Program, take charge of the space program for the alien race known as the Kerbals. You have access to an array of parts to assemble fully-functional spacecraft that flies (or doesn’t) based on realistic aerodynamic and orbital physics.",
  }
  game[15] = ginfo{
    Number: 16,
    Name: "ARK: Survival Evolved",
    Hrs: 26,
    Genre: "survival, PVE, PVP",
    Synopsis: "Stranded on the shores of a mysterious island, you must learn to survive. Use your cunning to kill or tame the primeval creatures roaming the land, and encounter other players to survive, dominate... and escape!",
  }
  game[16] = ginfo{
    Number: 17,
    Name: "Sid Meier's Civilization VI",
    Hrs: 25,
    Genre: "turn-based strategy, 4x",
    Synopsis: "a turn-based strategy game in which you attempt to build an empire to stand the test of time. Become Ruler of the World by establishing and leading a civilization from the Stone Age to the Information Age. Wage war, conduct diplomacy, advance your culture, and go head-to-head with history’s greatest leaders as you attempt to build the greatest civilization the world has ever known. Competing leaders will pursue their own agendas based on their historical traits as you race for one of five ways to achieve victory in the game.",
  }
  game[17] = ginfo{
    Number: 18,
    Name: "Left 4 dead 2",
    Hrs: 23,
    Genre: "survival, first person",
    Synopsis: "Set in the zombie apocalypse, Left 4 Dead 2 (L4D2) is the highly anticipated sequel to the award-winning Left 4 Dead. This co-operative action horror FPS takes you and your friends through the cities, swamps and cemeteries of the Deep South, from Savannah to New Orleans across five expansive campaigns. You'll play as one of four new survivors armed with a wide and devastating array of classic and upgraded weapons. In addition to firearms, you'll also get a chance to take out some aggression on infected with a variety of carnage-creating melee weapons, from chainsaws to axes and even the deadly frying pan.",
  }
  game[18] = ginfo{
    Number: 19,
    Name: "Planetary Annihilation: TITANS",
    Hrs: 22,
    Genre: "real-time strategy",
    Synopsis: "Wage war across entire solar systems with massive armies at your command. Annihilate enemy forces with world-shattering TITAN-class units, and demolish planets with massive super weapons!",
  }
  game[19] = ginfo{
    Number: 20,
    Name: "Final Fantasy IV",
    Hrs: 19.8,
    Genre: "RPG",
    Synopsis: "it is the fourth main installment of the Final Fantasy series. The game's story follows Cecil, a dark knight, as he tries to prevent the sorcerer Golbez from seizing powerful crystals and destroying the world",
  }
  game[20] = ginfo{
    Number: 21,
    Name: "South Park™: The Fractured But Whole™",
    Hrs: 19.3,
    Genre: "turn-based strategy, RPG",
    Synopsis: "Players will delve into the crime-ridden underbelly of South Park with Coon and Friends. This dedicated group of crime fighters was formed by Eric Cartman whose superhero alter-ego, The Coon, is half man, half raccoon. As the New Kid, players will join Mysterion, Toolshed, Human Kite and a host of others to battle the forces of evil while Coon strives to make his team the most beloved superheroes in history.",
  }
  game[21] = ginfo{
    Number: 22,
    Name: "Total War: Shogun 2",
    Hrs: 18.8,
    Genre: "turn-based strategy, real-time strategy",
    Synopsis: "In the darkest age of Japan, endless war leaves a country divided. It is the middle of the 16th Century in Feudal Japan. The country, once ruled by a unified government, is now split into many warring clans. Ten legendary warlords strive for supremacy as conspiracies and conflicts wither the empire. Only one will rise above all to win the heart of a nation as the new shogun...The others will die by his sword.Take on the role of one Daimyo, the clan leader, and use military engagements, economics and diplomacy to achieve the ultimate goal: re-unite Japan under his supreme command and become the new Shogun – the undisputed ruler of a pacified nation.",
  }
  game[22] = ginfo{
    Number: 23,
    Name: "Planetbase",
    Hrs: 17.4,
    Genre: "strategy, survival, simulation, city builder",
    Synopsis: "Guide a group of space settlers trying to establish a base on a remote planet. Grow food, collect energy, mine resources, survive disasters and build a self-sufficient colony in a harsh and unforgiving environment.",
  }
  game[23] = ginfo{
    Number: 24,
    Name: "Dead by Daylight",
    Hrs: 16.5,
    Genre: "horror, action, survival",
    Synopsis: "a multiplayer (4vs1) horror game where one player takes on the role of the savage Killer, and the other four players play as Survivors, trying to escape the Killer and avoid being caught and killed.",
  }
  game[24] = ginfo{
    Number: 25,
    Name: "Anno 2077",
    Hrs: 16,
    Genre: "real-time strategy, simulation, city builder",
    Synopsis: "2070. Our world has changed. The rising level of the ocean has harmed the coastal cities and climate change has made large stretches of land inhospitable. Master resources, diplomacy, and trade in the most comprehensive economic management system in the Anno series. Build your society of the future, colonize islands, and create sprawling megacities with multitudes of buildings, vehicles, and resources to manage. Engineer production chains such as Robot Factories, Oil Refineries, and Diamond Mines, and trade with a variety of goods and commodities.",
  }

 // inside program create variable for user to use for input
 var x string
 
 // inside main: print initial game info to include number and name for user  to make a selection
 for i:=0; i<24; i++{
   fmt.Println(game[i])
 }

 // inside main: print statement to explain how program works and how to work program
 fmt.Println("Above is a list of games played by me and the number associated with them in order from most played game down to least played. Please enter the game you would like to know more about by entering the associated number. You may also select 'random' if you need help picking somewhere to start.You may also type 'done' to end the program.")
 fmt.Scanln(&x)
 x = x-1
 // inside main: create loop that continues until user types "done"as selection or some itteration to complete program.
 for x != "Done" && x != "done" && x != "DONE" {
   // inside loop in main have loop to use rand func to get a random number as a string.
    for x == "random" || x == "Random" || x == "RANDOM" {
     X := randPick(x)
     fmt.Println("The random number you are given is",X)
     x = X
    }

    // inside itial loop until "done": after random, or if number selection made use if statements to get the appropriate info from selection to display 
    if x == "1" {
      fmt.Println("#",game1.Number,game1.Name,"hours played:",game1.Hrs,"\n","genre:",game1.Genre,"\n","synopsis:",game1.Synopsis)
    }else if x == "2" {
      fmt.Println("#",game2.Number,game2.Name,"hours played:",game2.Hrs,"\n","genre:",game2.Genre,"\n","synopsis:",game2.Synopsis)
    }else if x == "3" {
      fmt.Println("#",game3.Number,game3.Name,"hours played:",game3.Hrs,"\n","genre:",game3.Genre,"\n","synopsis:",game3.Synopsis)
    }else if x == "4" {
      fmt.Println("#",game4.Number,game4.Name,"hours played:",game4.Hrs,"\n","genre:",game4.Genre,"\n","synopsis:",game4.Synopsis)
    }else if x == "5" {
      fmt.Println("#",game5.Number,game5.Name,"hours played:",game5.Hrs,"\n","genre:",game5.Genre,"\n","synopsis:",game5.Synopsis)
    }else if x == "6" {
      fmt.Println("#",game6.Number,game6.Name,"hours played:",game6.Hrs,"\n","genre:",game6.Genre,"\n","synopsis:",game6.Synopsis)
    }else if x == "7" {
      fmt.Println("#",game7.Number,game7.Name,"hours played:",game7.Hrs,"\n","genre:",game7.Genre,"\n","synopsis:",game7.Synopsis)
    }else if x == "8" {
      fmt.Println("#",game8.Number,game8.Name,"hours played:",game8.Hrs,"\n","genre:",game8.Genre,"\n","synopsis:",game8.Synopsis)
    }else if x == "9" {
      fmt.Println("#",game9.Number,game9.Name,"hours played:",game9.Hrs,"\n","genre:",game9.Genre,"\n","synopsis:",game9.Synopsis)
    }else if x == "10" {
      fmt.Println("#",game10.Number,game10.Name,"hours played:",game10.Hrs,"\n","genre:",game10.Genre,"\n","synopsis:",game10.Synopsis)
    }else if x == "11" {
      fmt.Println("#",game11.Number,game11.Name,"hours played:",game11.Hrs,"\n","genre",game11.Genre,"synopsis:",game11.Synopsis)
    }else if x == "12" {
      fmt.Println("#",game12.Number,game12.Name,"hours played:",game12.Hrs,"\n","genre:",game12.Genre,"\n","synopsis:",game12.Synopsis)
    }else if x == "13" {
      fmt.Println("#",game13.Number,game13.Name,"hours played:",game13.Hrs,"\n","genre:",game13.Genre,"\n","synopsis:",game13.Synopsis)
    }else if x == "14" {
      fmt.Println("#",game14.Number,game14.Name,"hours played:",game14.Hrs,"\n","genre:",game14.Genre,"\n","synopsis:",game14.Synopsis)
    }else if x == "15" {
      fmt.Println("#",game15.Number,game15.Name,"hours played:",game15.Hrs,"\n","genre:",game15.Genre,"\n","synopsis:",game15.Synopsis)
    }else if x == "16" {
      fmt.Println("#",game16.Number,game16.Name,"hours played:",game16.Hrs,"\n","genre:",game16.Genre,"\n","synopsis:",game16.Synopsis)
    }else if x == "17" {
      fmt.Println("#",game17.Number,game17.Name,"hours played:",game17.Hrs,"\n","genre:",game17.Genre,"\n","synopsis:",game17.Synopsis)
    }else if x == "18" {
      fmt.Println("#",game18.Number,game18.Name,"hours played:",game18.Hrs,"\n","genre:",game18.Genre,"\n","synopsis:",game18.Synopsis)
    }else if x == "19" {
      fmt.Println("#",game19.Number,game19.Name,"hours played:",game19.Hrs,"\n","genre:",game19.Genre,"\n","synopsis:",game19.Synopsis)
    }else if x == "20" {
      fmt.Println("#",game20.Number,game20.Name,"hours played:",game20.Hrs,"\n","genre:",game20.Genre,"\n","synopsis:",game20.Synopsis)
    }else if x == "21" {
      fmt.Println("#",game21.Number,game21.Name,"hours played:",game21.Hrs,"\n","genre:",game21.Genre,"\n","synopsis:",game21.Synopsis)
    }else if x == "22" {
      fmt.Println("#",game22.Number,game22.Name,"hours played:",game22.Hrs,"\n","genre:",game22.Genre,"\n","synopsis:",game22.Synopsis)
    }else if x == "23" {
      fmt.Println("#",game23.Number,game23.Name,"hours played:",game23.Hrs,"\n","genre:",game23.Genre,"\n","synopsis:",game23.Synopsis)
    }else if x == "24" {
      fmt.Println("#",game24.Number,game24.Name,"hours played:",game24.Hrs,"\n","genre:",game24.Genre,"\n","synopsis:",game24.Synopsis)
    }else if x == "25" {
      fmt.Println("#",game25.Number,game25.Name,"hours played:",game25.Hrs,"\n","genre:",game25.Genre,"\n","synopsis:",game25.Synopsis)
    }
    
     // prompt for new selection after appropriate info displayed before repeating loop or ending program
    fmt.Println("Please enter the game you would like to know more about by entering the associated number. You may also select 'random' if you need help picking somewhere to start.You may also type 'done' to end the program.")
    fmt.Println(game1.Number,game1.Name,"\n",game2.Number,game2.Name,"\n",game3.Number,game3.Name,"\n",game4.Number,game4.Name,"\n",game5.Number,game5.Name,"\n",game6.Number,game6.Name,"\n",game7.Number,game7.Name,"\n",game8.Number,game8.Name,"\n",game9.Number,game9.Name,"\n",game10.Number,game10.Name,"\n",game11.Number,game11.Name,"\n",game12.Number,game12.Name,"\n",game13.Number,game13.Name,"\n",game14.Number,game14.Name,"\n",game15.Number,game15.Name,"\n",game16.Number,game16.Name,"\n",game17.Number,game17.Name,"\n",game18.Number,game18.Name,"\n",game19.Number,game19.Name,"\n",game20.Number,game20.Name,"\n",game21.Number,game21.Name,"\n",game22.Number,game22.Name,"\n",game23.Number,game23.Name,"\n",game24.Number,game24.Name,"\n",game25.Number,game25.Name)
    fmt.Scanln(&x)
    }

   fmt.Println("Thank you for using my program to learn about my top games. I hope one or more of these has sparked your interest and you get a chance to experience these great games.")  

}