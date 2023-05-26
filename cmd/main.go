package main

import (
	"github.com/anaregdesign/custom-skill-builder/service"
	"github.com/anaregdesign/custom-skill-builder/skill"
	"github.com/anaregdesign/custom-skill-example/process"
)

func main() {
	top10Skill := skill.NewSkill("content", "topWords", func(s string) ([]string, error) {
		return process.GetTopWord(s, 10)
	})
	book := skill.NewBook()
	book.Register("topwords", top10Skill.Flatten())

	svc := service.NewCustomSkillService(book)
	svc.Run()

}
