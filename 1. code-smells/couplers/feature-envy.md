# Feature Envy

## What It Looks Like

A method accesses the data of another object more than its own data. It occurs when a method in one class seems more interested in the fields or methods of another class rather than its own. It often indicates that the method should belong to the class it is envious of.


```
package main

import "fmt"

type address struct {
 day int
 fee int
}

type employee struct {
 id      int
 name    string
}

func (e employee) printEmployeeSalary(salary salary) {
 totalSalary := salary.day * salary.fee
 fmt.Printf("Employee %s will receive Rp. %d for %d days\n", e.name, totalSalary, salary.day)
}
```
Based on that example, there is wrong responsibility where employee calculate salary by itself.

## Why It Hurts

Feature envy is a code smell that points to a potential design issue, where methods are not residing in the most appropriate classes. Identifying and addressing feature envy through refactoring can lead to more maintainable, modular, and understandable code.

## How To Fix It

If a method clearly should be moved to another place, use [Move Method](.././../2.%20refactorings/move-method.md).
If only part of a method accesses the data of another object, use [Extract Method](.././../2.%20refactorings/extract-method.md).

## Refactor

```
package main

import "fmt"

type salary struct {
 day int
 fee int
}

func (s salary) totalSalary() int {
 return s.day * s.fee
}

type employee struct {
 id      int
 name    string
}

func (e employee) printEmployeeSalary(salary salary) {
 fmt.Printf("Employee %s will receive Rp. %d for %d days\n", e.name, salary.totalSalary(), salary.day)
}
```

## Payoff
- Less code duplication (if the data handling code is put in a central place).
- Better code organization (methods for handling data are next to the actual data).