DROP DATABASE IF EXISTS marvel_rivals;

CREATE DATABASE marvel_rivals;

\c marvel_rivals;

CREATE TABLE character (
	id SERIAL PRIMARY KEY,
	nickname VARCHAR (255) NOT NULL,
	first_name VARCHAR (255) NOT NULL,
	last_name VARCHAR (255) NOT NULL,
	role VARCHAR (255) NOT NULL,
	difficulty INTEGER NOT NULL,
	description TEXT NOT NULL,
	created_at TIMESTAMP DEFAULT NOW()
);

INSERT INTO character (nickname, first_name, last_name, role, difficulty, description)
	VALUES (
		'Hulk',
		'Bruce',
		'Banner',
		'Vanguard',
		4,
		'Brilliant scientist Dr. Bruce Banner has finally found a way to coexist with his monstrous alter ego, the Hulk. By accumulating gamma energy over multiple transformations, he can become a wise and strong Hero Hulk or a fierce and destructive Monster Hulk - a true force of fury on the battlefield!'
	);

INSERT INTO character (nickname, first_name, last_name, role, difficulty, description)
	VALUES (
		'The Punisher',
		'Frank',
		'Castle',
		'Duelist',
		1,
		'Expertly wielding a full arsenal of futuristic weapons, Frank Castle is a formidable one-man army. With a steadfast resolve to deliver justice to his enemies, The Punisher won''t cease in his mission until every last round is fired!'
	);

INSERT INTO character (nickname, first_name, last_name, role, difficulty, description)
	VALUES (
		'Storm',
		'Ororo',
		'Munroe',
		'Duelist',
		3,
		'An Omega-level Mutant ability to manipulate weather patterns makes Ororo Munroe a force to be reckoned with. Rain or shine, thunder or lightning, nature itself bends to the command of the Goddess of the Storm!'
	);

INSERT INTO character (nickname, first_name, last_name, role, difficulty, description)
	VALUES (
		'Loki',
		'Loki',
		'Laufeyson',
		'Strategist',
		4,
		'What greater thrill is there for a God of Mischief than to outsmart his foes? The cunning trickster Loki uses his illusions and shapeshifting abilities to weave in and out of combat, toying with enemies at every turn.'
	);

INSERT INTO character (nickname, first_name, last_name, role, difficulty, description)
	VALUES (
		'Doctor Strange',
		'Stephen',
		'Strange',
		'Vanguard',
		3,
		'As the Sorcerer Supreme, Doctor Stephen Strange gracefully wields ancient spells to turn the tide of even the most impossible battle. However, magic always comes at a cost, and each use of his arcane abilities gradually awakens the darkness within him.'
	);

INSERT INTO character (nickname, first_name, last_name, role, difficulty, description)
	VALUES (
		'Mantis',
		'Mantis',
		'',
		'Strategist',
		1,
		'Mantis uses her impressive mental abilities and her penchant for plant control to anchor any team she fights alongside. Her powers tap into a limitless flow of life energy, gently nourishing everything she touches.'
	);

INSERT INTO character (nickname, first_name, last_name, role, difficulty, description)
	VALUES (
		'Rocket Raccoon',
		'Rocket',
		'Raccoon',
		'Strategist',
		1,
		'Rocket may not look like a tech genius or an expert tactician, but anyone who''s ever made his hit list has quickly regretted underestimating him. This savvy space soldier is equally eager to boost his teammates and to collect bounties on his foes.'
	);

INSERT INTO character (nickname, first_name, last_name, role, difficulty, description)
	VALUES (
		'Hela',
		'Hela',
		'',
		'Duelist',
		3,
		'As the Goddess of Death, Hela wields supreme control over the fallen souls residing in Hel. With a haunting whisper and a murder of crows, the queen of the underworld gracefully reaps the souls of her enemies without an ounce of mercy.'
	);

INSERT INTO character (nickname, first_name, last_name, role, difficulty, description)
	VALUES (
		'Black Panther',
		'T''Challa',
		'',
		'Duelist',
		4,
		'T''Challa, King of Wakanda, wields the perfect blend of the cutting-edge Vibranium technology and ancestral power drawn from the Panther God, Bast. The Black Panther bides his time until elegantly infiltrating enemy lines and commencing his hunt.'
	);

INSERT INTO character (nickname, first_name, last_name, role, difficulty, description)
	VALUES (
		'Groot',
		'Groot',
		'',
		'Vanguard',
		1,
		'A flora colossus from Planet X, the alien known as Groot exhibits enhanced vitality and the ability to manipulate all forms of vegetation. As sturdy as a towering tree, Groot forges his own way, serving as the team''s silent but reliable pathfinder.'
	);

INSERT INTO character (nickname, first_name, last_name, role, difficulty, description)
	VALUES (
		'Magik',
		'Illyana',
		'Rasputin',
		'Duelist',
		3,
		'Trained in the dark arts and wielding her mighty Soulsword, Magik leaps through portals to navigate the arena with ease. Once Illyana transforms into the demonic Darkchild, all who dare stand against her will fall before her merciless blade.'
	);

INSERT INTO character (nickname, first_name, last_name, role, difficulty, description)
	VALUES (
		'Luna Snow',
		'Seol',
		'Hee',
		'Strategist',
		2,
		'Equal parts pop star and Super hero, Luna Snow puts on a dazzling show with both her light and dark ice powers. The arena is her stage, where Seol Hee and her team orchestrate spectacular displays that earn her an ever-increasing number of fans and wins.'
	);

INSERT INTO character (nickname, first_name, last_name, role, difficulty, description)
	VALUES (
		'Iron Man',
		'Anthony "Tony"',
		'Stark',
		'Duelist',
		2,
		'Armed with superior intellect and a nanotech battlesuit of his own design, Tony Stark stands alongside gods as the Invincible Iron Man. His state of the art armor turns any battlefield into his personal playground, allowing him to steal the spotlight he so desperately desires.'
	);

INSERT INTO character (nickname, first_name, last_name, role, difficulty, description)
	VALUES (
		'Spider-Man',
		'Peter',
		'Parker',
		'Duelist',
		5,
		'Swinging around the arena on his signature weblines, your friendly neighborhood Spider-Man, AKA Peter Parker, catches his rivals by surprise with sneaky, sticky bursts of webbing and unexpected attacks from above. Look out... here comes the Spider-Man!'
	);

INSERT INTO character (nickname, first_name, last_name, role, difficulty, description)
	VALUES (
		'Magneto',
		'Max',
		'Eisenhardt',
		'Vanguard',
		2,
		'The Master of Magnestism bends even the strongest metal to his whims, shielding his allies and striking at his foes. Whether he calls himself Max Eisenhardt, Erik Lehnsherr, or simply Magneto, the hardships this warrior has endured have made him as unbreakable as the steel he brandishes.'
	);

INSERT INTO character (nickname, first_name, last_name, role, difficulty, description)
	VALUES (
		'Scarlet Witch',
		'Wanda',
		'Maximoff',
		'Duelist',
		1,
		'Wanda Maximoff is adept at harnessing formidable chaos magic, casting hexes with the power to twist and reshape reality itself. Energy, space, and matter are mere playthings in the hands of Scarlet Witch!'
	);

INSERT INTO character (nickname, first_name, last_name, role, difficulty, description)
	VALUES (
		'Peni Parker',
		'Peni',
		'Parker',
		'Vanguard',
		3,
		'Peni Parker may be young, but she bravely stands on the frontlines to protect the Web of Life and Destiny. Together, this teen prodigy and her state-of-the-art mech, the sensational SP//dr, make for the most thrilling duo on the battlefield!'
	);

INSERT INTO character (nickname, first_name, last_name, role, difficulty, description)
	VALUES (
		'Star-Lord',
		'Peter',
		'Quill',
		'Duelist',
		2,
		'Peter Quill lives to dazzle his foes on the battlefield with his signature swagger. As his element guns paint arcs of devastation, his acrobatic moves sail through the sky with unrivaled style. With performances this spectacular, it''s no wonder that Star-Lord is so legendary!'
	);

INSERT INTO character (nickname, first_name, last_name, role, difficulty, description)
	VALUES (
		'Namor',
		'Namor',
		'',
		'Duelist',
		2,
		'The unrivaled King of the Seas, Namor surfs into battle on a mighty wave with an army of fierce aquatic creatures in his wake. When ancient horns of war blare, devastation soon follows as deadly waters engulf the arena.'
	);

INSERT INTO character (nickname, first_name, last_name, role, difficulty, description)
	VALUES (
		'Venom',
		'Edward "Eddie"',
		'Brock',
		'Vanguard',
		3,
		'Using his symbiote- enhanced body as the perfect living weapon, Eddie Brock and his alien ally stand ever-ready to unleash vicious attacks upon anyone he deems an enemy. Those ensnared by Venom''s tentacles have no choice but to surrender to this insatiable predator.'
	);

INSERT INTO character (nickname, first_name, last_name, role, difficulty, description)
	VALUES (
		'Adam Warlock',
		'Adam',
		'',
		'Strategist',
		2,
		'The genetically- engineered Adam Warlock wields mighty Quantum Magic, allowing him to connect and heal souls with a gentle touch. When the time comes for his allies to unite, Warlock emerges as the unwavering epicenter of cosmic justice!'
	);

INSERT INTO character (nickname, first_name, last_name, role, difficulty, description)
	VALUES (
		'Thor',
		'Thor',
		'Odinson',
		'Vanguard',
		2,
		'The son of Odin taps into his divine power to call forth thunder and lightning, raining down relentless fury upon his enemies. With his mighty hammer Mjolnir in hand, Thor effortlessly asserts his dominance on the field of combat.'
	);

INSERT INTO character (nickname, first_name, last_name, role, difficulty, description)
	VALUES (
		'Jeff The Land Shark',
		'Jeff',
		'',
		'Strategist',
		2,
		'Most land sharks are vicious creatures of the deep... but not Jeff! This adorable and mischievous little land shark brings splashes of joy and healing to every battle. But if the tide turns, Jeff can morph into a voracious beast, swallowing an army of foes in one giant gulp!'
	);

INSERT INTO character (nickname, first_name, last_name, role, difficulty, description)
	VALUES (
		'Winter Soldier',
		'James "Bucky"',
		'Barnes',
		'Duelist',
		2,
		'Terrifying experiments turned him into a brainwashed assassin, but now james "Bucky" Barnes is in control of his own fate once again. With his enhanced mechanical arm, the Winter Soldier is primed to deliver earth- shattering blows to any foe in his path!'
	);

INSERT INTO character (nickname, first_name, last_name, role, difficulty, description)
	VALUES (
		'Captain America',
		'Steven "Steve"',
		'Rogers',
		'Vanguard',
		2,
		'Enhanced by the Super- Soldier Serum, Steven "Steve" Rogers uses his Vibranium shield and extensive combat training to confront any threat to justice. When Captain America rallies his troops, a wave of courage sweeps across the battlefield!'
	);