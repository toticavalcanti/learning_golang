package main

import "fmt"

func main() {
    var test_grade_1, test_grade_2, test_grade_3, test_grade_4 float64
    fmt.Scan(&test_grade_1, &test_grade_2, &test_grade_3, &test_grade_4)
    average := (test_grade_1*2 + test_grade_2*3 + test_grade_3*4 + test_grade_4*1) / (2 + 3 + 4 + 1)
    fmt.Printf("Media: %.1f\n", average)
    if average >= 7 {
        fmt.Println("Aluno aprovado.")
    } else if average < 5 {
        fmt.Println("Aluno reprovado.")
    } else {
        fmt.Println("Aluno em exame.")
        var s float64
        fmt.Scan(&s)
        fmt.Printf("Nota do exame: %.1f\n", s)
        final_average := (average + s) / 2
        if final_average >= 5 {
            fmt.Println("Aluno aprovado.")
        } else if final_average <= 4.9 {
            fmt.Println("Aluno reprovado.")
        }
        fmt.Printf("Media final: %.1f\n", final_average)
    }
}