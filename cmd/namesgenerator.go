package cmd

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

var (
	names = []string{
		// Ada Lovelace - English mathematician and writer, known for her work on Charles Babbage's early mechanical general-purpose computer, the Analytical Engine
		"lovelace",
		// Alan Turing - British pioneering computer scientist, cryptanalyst, and father of theoretical computer science and artificial intelligence
		"turing",
		// Grace Hopper - American computer scientist, developed the first compiler for a computer programming language
		"hopper",
		// Albert Einstein - German-born theoretical physicist, developed the theory of relativity
		"einstein",
		// Marie Curie - Polish and naturalized-French physicist and chemist who conducted pioneering research on radioactivity
		"curie",
		// DaVinci - Italian polymath of the Renaissance whose areas of interest included invention, painting, sculpting, architecture, science, music, mathematics, engineering, literature, anatomy, geology, astronomy, botany, writing, history, and cartography
		"davinci",
		// Nikola Tesla - Serbian-American inventor, electrical engineer, mechanical engineer, and futurist
		"tesla",
		// Richard Feynman - American theoretical physicist known for his work in quantum mechanics and particle physics
		"feynman",
		// Linus Torvalds - Finnish-American software engineer, creator of the Linux kernel
		"torvalds",
		// Richard Stallman - American free software movement activist and programmer
		"stallman",
		// Dennis Ritchie - American computer scientist, created the C programming language
		"ritchie",
		// Donald Knuth - American computer scientist, mathematician, and professor emeritus at Stanford University
		"knuth",
		// Edsger Dijkstra - Dutch computer scientist and pioneer in many research areas of computing science
		"dijkstra",
		// Steve Wozniak - American electronics engineer, cofounder of Apple Computer Inc.
		"wozniak",
		// Tim Berners-Lee - English computer scientist, best known as the inventor of the World Wide Web
		"bernerslee",
		// James Gosling - Canadian computer scientist, best known as the father of the Java programming language
		"gosling",
		// Yukihiro Matsumoto - Japanese computer scientist and software programmer best known as the chief designer of the Ruby programming language
		"matsumoto",
		// Ken Thompson - American pioneer of computer science, co-creator of the Unix operating system
		"thompson",
		// Gordon Moore - American businessman, engineer, and co-founder of Intel Corporation
		"moore",
		// Niels Bohr - Danish physicist who made foundational contributions to understanding atomic structure and quantum theory
		"bohr",
		// Stephen Hawking - English theoretical physicist, cosmologist, and author
		"hawking",
		// Ramajuan - Indian mathematician who lived during the British Rule in India. Though he had almost no formal training in pure mathematics, he made substantial contributions to mathematical analysis, number theory, infinite series, and continued fractions, including solutions to mathematical problems considered to be unsolvable.
		"ramajuan",
		// Dawkins - British ethologist, evolutionary biologist, and author. He is an emeritus fellow of New College, Oxford, and was the University of Oxford's Professor for Public Understanding of Science from 1995 until 2008.
		"dawkins",
		// Sagan - American astronomer, planetary scientist, cosmologist, astrophysicist, astrobiologist, author, and science communicator. His best known scientific contribution is research on extraterrestrial life, including experimental demonstration of the production of amino acids from basic chemicals by radiation.
		"sagan",
		// Dawin - English naturalist, geologist and biologist, best known for his contributions to the science of evolution.
		"darwin",
		// Newton - English mathematician, physicist, astronomer, theologian, and author who is widely recognised as one of the most influential scientists of all time and as a key figure in the scientific revolution.
		"newton",
		// Galileo - Italian astronomer, physicist and engineer, sometimes described as a polymath, from Pisa. Galileo has been called the "father of observational astronomy", the "father of modern physics", the "father of the scientific method", and the "father of modern science".
		"galileo",
		// diffie - American cryptographer and one of the pioneers of public-key cryptography.
		"diffie",
		// hellman - American cryptographer and one of the pioneers of public-key cryptography.
		"hellman",
		// dijkstra - Dutch computer scientist, programmer, software engineer, systems scientist, science essayist, and pioneer in computing science.
		"dijkstra",
		// knuth - American computer scientist, mathematician, and professor emeritus at Stanford University.
		"knuth",
		// Feynman - American theoretical physicist known for his work in quantum mechanics and particle physics.
		"feynman",
		// oppenheimer - American theoretical physicist and the head of the Los Alamos Laboratory during World War II.
		"oppenheimer",
		// Morse - American painter and inventor of the single-wire telegraph and Morse code.
		"morse",
		// Neumann - Hungarian-American mathematician, physicist, computer scientist, engineer and polymath.
		"neumann",
		// Rob Pike - Canadian programmer and author.
		"pike",
		// Brian Kernighan - Canadian computer scientist.
		"kernighan",
	}
)

func getRandomName() (string, error) {
	max := big.NewInt(int64(len(names)))
	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		return "", fmt.Errorf("failed to generate random names: %w", err)
	}
	return names[n.Int64()], nil
}
