<!-- insert
---
title: "💎 Readme In Static Site (RISS)"
date: 2021-08-21T08:15:54
---
{{< github_badge >}}
end_insert -->
<!-- remove -->
# 💎 Readme In Static Site (RISS)
<!-- end_remove -->

Insert your GitHub readme in your static site and apply transformations. For instance, you can read this [readme on Github](https://github.com/cljoly/readme-in-static-site/blob/main/README.md) and [on my website](https://joly.pw/readme-in-static-site).

### Why?

The GitHub Readme of your repo needs to efficiently describe your project to GitHub’s visitor. But featuring your project on your website allows you to (among other things):
- have more control on the theme and layout,
- insert scripts that GitHub would prohibit (like [asciinema](#replace-asciinema-image)),
- have your project’s homepage independant from your hosting platform.

Chances are that for small projects the page about your project is very similar to the GitHub readme. Don’t duplicate efforts, just describe the differences!

### Run it (nothing to install)

To try it with [Hugo][hugo] or [Zola][zola], run the following in your static-site sources:
```sh
wget https://joly.pw/riss.awk
awk -f riss.awk /path/to/my-project/readme.md > content/my-project.md
```

If you don’t use Hugo or Zola, no problem! It should also work with any markdown-based static-site generator. Just put the markdown file where it makes sense for that generator.

To automatically update these files in your static-site sources, see [Automate with GitHub Actions](#automate-with-github-actions) below.

## Example

### Add a front matter

Most static site generators require a “[frontmatter](https://gohugo.io/getting-started/configuration/#configure-front-matter)” at the beginning of a markdown file to attach some metadata. But you don’t want to add this on your GitHub readme! Let’s hide this on GitHub and have it in the script’s output.

In you .md file on GitHub, put:

    <!-- insert
    ---
    title: "Readme In Static Site (RISS)"
    date: 2021-08-21T10:15:54
    ---
    end_insert -->
    <!-- remove -->
    
    # Readme In Static Site (RISS)
    <!-- end_remove -->

The output of the script will be:

    ---
    title: "Readme In Static Site (RISS)"
    date: 2021-08-21T10:15:54
    ---

and this piece of yaml will be hidden on GitHub!

### Replace Asciinema Image

You can’t embed the asciinema player on GitHub for security reasons. So the [asciinema documentation](https://asciinema.org/docs/embedding) suggests to use an image there and to link it to a webpage with the player. But on your own website, you can embed this player.

In your .md file, put:

    <!-- remove -->
    [![Finding the repositories with “telescope” in their name, with the README in the panel on the right](https://asciinema.org/a/427156.svg)](https://asciinema.org/a/427156)
    <!-- end_remove -->
    <!-- insert
    <asciinema-player src="./telescope.cast" poster="npt:0:04"></asciinema-player>
    end_insert -->

The output will contain only the asciinema player:

    <asciinema-player src="./telescope.cast" poster="npt:0:04"></asciinema-player>

*Note*: you also need to add the JS/CSS files of the asciinema player somewhere in your theme.

### More

See the [input file (typically on GitHub)](https://github.com/cljoly/readme-in-static-site/blob/main/test.md) and the [output of the script](https://github.com/cljoly/readme-in-static-site/blob/main/test_output.md). You can find another real word [readme](https://github.com/cljoly/telescope-repo.nvim/blob/master/README.md) converted to a [webpage](https://joly.pw/telescope-repo-nvim/) (this gives an other example of asciinema conversion using a Hugo shortcode).

## Transformations Reference

The transformations are driven by html comments, so that you can have different results when comments are ignored (e.g. in your GitHub readme) and after executing the script on your markdown file.

### Escaping

**It is important that your comment starts at the beginning of the line:** spaces are used for escaping, meaning that if the comment has spaces at the beginning of the line, it is ignored.

So this is escaped
```html
    <!-- insert
```
but this is not

    <!-- insert

### Insertion

Use these two lines for text you want to have in your output, but not in the raw .md file.

    <!-- insert
    end_insert -->

### Remove

Use these two comments for text you want to have in your raw .md file, but not in the output

    <!-- remove -->
    <!-- end_remove -->

## Automate with GitHub Actions

You can automatically update the markdown file in the sources of your site with GitHub Actions. You can put this very simple workflow in `.github/workflows/readme.yml`:

```yaml
name: Update README files

on:
  schedule:
    - cron: '30 */2 * * *'
  push:
    branches:
    - master
  # To run this workflow manually from GitHub GUI
  workflow_dispatch:

jobs:
  build:
    runs-on: ubuntu-latest
    steps:
    - name: Check out the repo
      uses: actions/checkout@v2
    - name: Get the latest READMEs
      run: make readme-update
    - name: Commit and push if there are changes
      run: |-
        git diff
        git config --global user.email "bot@example.com"
        git config --global user.name "bot"
        git diff --quiet || (git add -u && git commit -m "Update READMEs")
        git push
```

and then your `Makefile` may contain something like:

```make
readme-update:
	curl https://raw.githubusercontent.com/cljoly/readme-in-static-site/main/README.md | awk -f riss.awk >content/readme-in-static-site.md
```

## Contributions are welcome!

Feel free to open an issue to discuss something or to send a PR.

[hugo]: https://gohugo.io/
[zola]: https://www.getzola.org/
