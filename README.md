# Roman Numerals Conversion - Written in go

## Prerequisits

GoLang [https://medium.com/@betakuang/setup-go-development-environment-with-vs-code-and-wsl-on-windows-62bd4625c6a7]

## Getting started

Run the build:

```bash
./run.sh
```

Run the tests:

```bash
./test.sh
```

## The problem

Convert Roman numerals into decimal using the following table of symbol values:-

I = 1
V = 5
X = 10
L = 50
C = 100
D = 500
M = 1000

Generally, symbols are placed in order of value, starting with the largest values. When smaller values precede larger values, the smaller values are subtracted from the larger values, and the result is added to the total. However, you can’t write numerals like “IM” for 999, there are some additional rules:

• A number written in Arabic numerals can be broken into digits. For example, 1903 is composed of 1 (one thousand), 9 (nine hundreds), 0 (zero tens), and 3 (three units). To write the Roman numeral, each of the nonzero digits should be treated separately. In the above example, 1,000 = M, 900 = CM, and 3 = III. Therefore, • 1903 = MCMIII.

• The symbols “I”, “X”, “C”, and “M” can be repeated three times in succession, but no more. (They may appear more than three times if they appear non-sequentially, such as XXXIX.) “D”, “L”, and “V” can never be repeated. • “I” can be subtracted from “V” and “X” only. “X” can be subtracted from “L” and “C” only. “C” can be subtracted from “D” and “M” only. “V”, “L”, and “D” can never be subtracted. • Only one small-value symbol may be subtracted from any large-value symbol.

