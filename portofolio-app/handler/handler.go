package handler

import (
	"bufio"
	"fmt" 
	"os"
	"strconv"
	"strings"

	"portofolio-app/models"
)


func RunApp() {
	for {
		fmt.Println("=== Menu Portofolio ===")
		fmt.Println("1. Tambah Proyek")
		fmt.Println("2. Lihat Semua Proyek")
		fmt.Println("3. Keluar")
		fmt.Print("Pilih Menu: ")

		choice := readInput()

		switch choice {
		case "1":
			addProjectHandler()
		case "2":
			viewProjectHandler()
		case "3":
			fmt.Println("Sampai Jumpa!")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func addProjectHandler() {
	fmt.Println("\n--- Tambah Proyek ---")

	fmt.Print("Nama Proyek: ")
	name := readInput()

	fmt.Print("Deskripsi: ")
	description := readInput()

	fmt.Print("Skill (pisahkan dengan koma): ")
	skillsInput := readInput()
	skills := strings.Split(skillsInput, ",")

	fmt.Print("Tahun: ")
	yearInput := readInput()
	year, _ := strconv.Atoi(yearInput)

	newProject := models.Project{
		ID:          len(models.Projects) + 1,
		Name:        name,
		Description: description,
		Skills:      skills,
		Year:        year,
	}

	models.AddProject(newProject)
	fmt.Println("Proyek Berhasil Ditambahkan!\n")
}

func viewProjectHandler() {
	fmt.Println("\n--- Daftar proyek ---")
	projects := models.GetAllProjects()

	if len(projects) == 0 {
		fmt.Println("Belum ada proyek.")
		return
	}

	for _, p := range projects {
		fmt.Printf("ID: %d\nNama: %s\nDeskripsi: %s\nSkills: %v\nTahun: %d\n\n", p.ID, p.Name, p.Description, p.Skills, p.Year)

	}
}

func readInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}
