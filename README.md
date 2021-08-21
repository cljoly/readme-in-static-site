<!-- insert
title: "Readme In Static Site (RISS)"
date: 2021-01-17T20:44:12Z
end_insert -->
<!-- remove -->
# Readme In Static Site (RISS)
<!-- end_remove -->

Insert your github readme in your static site (hugo, zola…) with small transformations, for instance to add a [frontmatter](https://gohugo.io/getting-started/configuration/#configure-front-matter).

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

### [Input (typically on GitHub)](./test.md)

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

### [Output (typically in zola or hugo)](./test_output.md)
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
