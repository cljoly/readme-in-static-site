# readme-static-site

Insert your github readme in your static site (hugo, zola…) with small transformations.

## Codes

### Insertion

Use these two lines for text you want to have in your output, but not in the raw .md file.
```html
<!-- insert
```

```html
end_insert -->
```

### Remove

Use these two comments for text you want to have in your raw .md file, but not in the output
```html
<!-- remove -->
```
```html
<!-- end_remove -->
```

## Example

### Input

```markdown
<!-- insert
---
title: "My Page Title"
date: "2222-11-11"
---
end_insert -->

<!-- remove -->
Present in the .md, removed in the output
<!-- end_remove -->

**kept**
Present in both the .md and the output

<!-- insert
**inserted**
only in the output
end_insert -->

<!-- insert
some complicated html
<script id="asciicast-427156" src="https://asciinema.org/a/427156.js" async></script>
end_insert -->
```

### Output
```
---
title: "My Page Title"
date: "2222-11-11"
---


**kept**
Present in both the .md and the output

**inserted**
only in the output

some complicated html
<script id="asciicast-427156" src="https://asciinema.org/a/427156.js" async></script>
```
