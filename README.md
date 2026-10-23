# Math Skills

This program reads a list of numbers from a file and calculates:

- Average  
- Median  
- Variance  
- Standard Deviation  

All results are **rounded integers**, as required by the subject.

---

## 📥 Input Format

The program expects a text file where **each line contains one integer**.

Example:

```
189
113
121
114
145
110
```

---

## ▶️ How to Run the Student Program

Run the program using:

```bash
go run math-skills.go data.txt
```

Where **data.txt** is the input file created by the testing script.

---

## 🧪 How to Test This Program (Required by Auditors)

### 1. Run the official testing script

You can do this in two ways:

```bash
./bin/math-skills
```

or:

```bash
./run.sh math-skills
```

The script will:

- Generate a new `data.txt` file with random values  
- Print the **reference** results for:
  - Average  
  - Median  
  - Variance  
  - Standard Deviation  

---

## ▶️ Step 2 — Run the Student Program

After the script generates `data.txt`, run the student's solution:

```bash
go run math-skills.go data.txt
```

The output will look like this:

```
Average: 35
Median: 4
Variance: 5
Standard Deviation: 65
```

(These values are only examples.)

---

## ✅ What the Auditor Must Check

### ✔️ Format Check  
Program output must match this format exactly:

```
Average: X
Median: Y
Variance: Z
Standard Deviation: W
```

### ✔️ Rounding Check  
All printed values must be **rounded integers**.

### ✔️ Correctness Check  
Values produced by the student program must **match exactly** the values printed by the official script.

---

## 🔄 Step 3 — Repeat Testing

Repeat the same procedure **3 more times**:

1. Run the testing script again:
   ```bash
   ./bin/math-skills
   ```
2. Run the student program:
   ```bash
   go run math-skills.go data.txt
   ```
3. Compare the results.

The values must match **every time**.

---

## ✔️ Notes

- Empty lines in the input file are ignored.  
- Non-numeric values produce an error.  
- Median works for even and odd sets.  
- Standard deviation is computed using `math.Sqrt()` on the variance.

---
