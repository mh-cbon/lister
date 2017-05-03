# Changelog - lister

### 0.0.1

__Changes__

- __updated__
  - Filters: the method now accepts multiple input filters func
- __new__
  - Filters: Automatically generate pre defined filters for basic types
  - Map: Added a new function to map a slice
- __dependency__
  - update astutil



__Contributors__

- mh-cbon

Released by mh-cbon, Wed 03 May 2017 -
[see the diff](https://github.com/mh-cbon/lister/compare/0.0.1-beta5...0.0.1#diff)
______________

### 0.0.1-beta5

__Changes__

- README
- add new methods Contains, Empty, Last, First
- add new methods
  - __Contains__(s type) bool: returns true if the value s exists in t.
  - __Empty__() bool: returns true if t is empty.
  - __Last__() type: returns first or defaut value
  - __First__() type: returns last or defaut value.





__Contributors__

- mh-cbon

Released by mh-cbon, Mon 01 May 2017 -
[see the diff](https://github.com/mh-cbon/lister/compare/0.0.1-beta4...0.0.1-beta5#diff)
______________

### 0.0.1-beta4

__Changes__

- add __Filter__(f func(type) bool) []type: returns a new array with items matching f.

__Contributors__

- mh-cbon

Released by mh-cbon, Sun 30 Apr 2017 -
[see the diff](https://github.com/mh-cbon/lister/compare/0.0.1-beta3...0.0.1-beta4#diff)
______________

### 0.0.1-beta3

__Changes__

- Changed type implementation: do not use type alias `type xx []yy`, prefer an embeded property of `[]y`
- added test

__Contributors__

- mh-cbon

Released by mh-cbon, Sun 30 Apr 2017 -
[see the diff](https://github.com/mh-cbon/lister/compare/0.0.1-beta2...0.0.1-beta3#diff)
______________

### 0.0.1-beta2

__Changes__

- fix constructor comment to use unpointed type
- glide update

__Contributors__

- mh-cbon

Released by mh-cbon, Sun 30 Apr 2017 -
[see the diff](https://github.com/mh-cbon/lister/compare/0.0.1-beta1...0.0.1-beta2#diff)
______________

### 0.0.1-beta1

__Changes__

- dep: add glide support
- refactoring to use astutil package

__Contributors__

- mh-cbon

Released by mh-cbon, Sat 29 Apr 2017 -
[see the diff](https://github.com/mh-cbon/lister/compare/0.0.1-beta...0.0.1-beta1#diff)
______________

### 0.0.1-beta

__Changes__

- Project initialization.

__Contributors__

- mh-cbon

Released by mh-cbon, Sat 29 Apr 2017 -
[see the diff](https://github.com/mh-cbon/lister/compare/069a44027103ddd015e5fca0f286d5a0eaad3464...0.0.1-beta#diff)
______________


