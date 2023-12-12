# Feature Envy

## What It Looks Like

A method accesses the data of another object more than its own data. It occurs when a method in one class seems more interested in the fields or methods of another class rather than its own. It often indicates that the method should belong to the class it is envious of. Feature Envy often shows up as a misplaced responsibility.


```
package main

type Phone struct {
	unformattedNumber string
}

func NewPhone(unformattedNumber string) *Phone {
	return &Phone{unformattedNumber: unformattedNumber}
}

func (p *Phone) getAreaCode() string {
	return p.unformattedNumber[0:3]
}

func (p *Phone) getPrefix() string {
	return p.unformattedNumber[3:6]
}

func (p *Phone) getNumber() string {
	return p.unformattedNumber[6:10]
}

type Customer struct {
	mobilePhone *Phone
}

func (c *Customer) GetMobilePhoneNumber() string {
	if c.mobilePhone != nil {
		return fmt.Sprintf("(%s) %s-%s", c.mobilePhone.getAreaCode(), c.mobilePhone.getPrefix(), c.mobilePhone.getNumber())
	}
	return ""
}
```
Based on that example, there is wrong responsibility where customer formats the phone number by itself.

## Why It Hurts

Feature envy is a code smell that points to a potential design issue, where methods are not residing in the most appropriate classes. Identifying and addressing feature envy through refactoring can lead to more maintainable, modular, and understandable code.

## How To Fix It

If a method clearly should be moved to another place, use [Move Method](.././../2.%20refactorings/move-method.md).
If only part of a method accesses the data of another object, use [Extract Method](.././../2.%20refactorings/extract-method.md).

## Refactor

```
package main

type Phone struct {
	unformattedNumber string
}

func NewPhone(unformattedNumber string) *Phone {
	return &Phone{unformattedNumber: unformattedNumber}
}

func (p *Phone) getAreaCode() string {
	return p.unformattedNumber[0:3]
}

func (p *Phone) getPrefix() string {
	return p.unformattedNumber[3:6]
}

func (p *Phone) getNumber() string {
	return p.unformattedNumber[6:10]
}

func (p *Phone) ToFormattedString() string {
	return fmt.Sprintf("(%s) %s-%s", p.getAreaCode(), p.getPrefix(), p.getNumber())
}

type Customer struct {
	mobilePhone *Phone
}

func (c *Customer) GetMobilePhoneNumber() string {
	if c.mobilePhone != nil {
		return c.mobilePhone.ToFormattedString()
	}
	return ""
}
```
Now Customer relies on Phone to do the formatting.

## Payoff
- Less code duplication (if the data handling code is put in a central place).
- Better code organization (methods for handling data are next to the actual data).