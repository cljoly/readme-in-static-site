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

See the [input (typically on GitHub)](https://github.com/cljoly/readme-in-static-site/blob/main/test.md) and the [output of the script](https://github.com/cljoly/readme-in-static-site/blob/main/test_output.md)
