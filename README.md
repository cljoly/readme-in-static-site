<!-- insert
---
title: "README In Static Site (RISS)"
date: 2021-08-21T08:15:54
description: "💎 Transform and insert your GitHub readme in your static site"
tags:
- Open Source
- Tools
- Static Site
- Hugo
- HackerNewsHomepage
repo_url: "https://codeberg.org/cljoly/readme-in-static-site"
---
{{< github_badge >}}
{{< rawhtml >}}
<div class="badges">
{{< /rawhtml >}}
end_insert -->
<!-- Powered by https://cj.rs/riss -->

<!-- remove -->
# 💎 README In Static Site (RISS)
<!-- end_remove -->

[![Get it on codeberg](https://raw.githubusercontent.com/cljoly/readme-in-static-site/refs/heads/moved/get-it-on-codeberg.svg)](https://codeberg.org/cljoly/readme-in-static-site)

[![GitHub code size in bytes](https://img.shields.io/github/size/cljoly/readme-in-static-site/riss.awk?color=purple)](https://github.com/cljoly/readme-in-static-site/blob/main/riss.awk) ![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/cljoly/readme-in-static-site?color=darkgreen&sort=semver) [![CI](https://codeberg.org/cljoly/readme-in-static-site/actions/workflows/checks.yml/badge.svg)](https://github.com/cljoly/readme-in-static-site/actions/workflows/checks.yml) [![](https://img.shields.io/badge/powered%20by-riss-lightgrey)](https://cj.rs/riss)

<!-- remove -->
![RISS in action](https://user-images.githubusercontent.com/7347374/134990638-d1ebb3ba-89db-4bbf-b00f-dc150e3b35fc.png)
<!-- end_remove -->
<!-- insert
![RISS in action](/blog/putting-readmes-on-your-static-site/riss-in-action.png)
end_insert -->

<!-- insert
{{< rawhtml >}}
</div>
{{< /rawhtml >}}
end_insert -->

This [fast][] [script][] allows you to insert your GitHub README in your static site and apply transformations. For instance, you can read this [README on GitHub](https://github.com/cljoly/readme-in-static-site/blob/main/README.md) and [on my website](https://cj.rs/readme-in-static-site).

### Why?

The GitHub README of your repo needs to efficiently describe your project to GitHub’s visitor. But featuring your project on your website allows you to (among other things):
* have more control on the theme and layout,
* insert scripts that GitHub would prohibit (like [asciinema](#replace-asciinema-image)),
* have your project’s homepage independent of your hosting platform, if you wish to change at some point.

Chances are that for small projects, the page about your project is very similar to the GitHub README. Don’t duplicate efforts, describe the differences! This [has](https://github.com/vhodges/stitcherd) [been](https://dev.to/lornasw93/github-readme-on-portfolio-project-page-51i8) a [long-awaited](https://richjenks.com/github-pages-from-readme/) [feature](https://medium.com/@bolajiayodeji/how-to-convert-github-markdown-files-to-a-simple-website-b08602d05e1c), in [one](https://news.ycombinator.com/item?id=29305990) [form](https://stackoverflow.com/q/48919200/4253785) or [another](https://stackoverflow.com/a/69296054/4253785).

See this [blog post](https://cj.rs/blog/putting-readmes-on-your-static-site/) for more details.

## Testimonials

[![](https://img.shields.io/badge/dynamic/json?color=Orange&label=HackerNews&query=%24.score&url=https%3A%2F%2Fhacker-news.firebaseio.com%2Fv0%2Fitem%2F29304376.json&logo=ycombinator&color=orange)](https://news.ycombinator.com/item?id=29304376) [![](https://img.shields.io/badge/dynamic/json?color=green&label=Lobsters&query=%24.score&url=https%3A%2F%2Flobste.rs%2Fs%2Fa4jzvv%2Freadme_static_site_riss.json&logo=data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAABAAAAAQCAYAAAAf8/9hAAAAAXNSR0IArs4c6QAAAAZiS0dEAL0ALQAtZF7+HAAAAAlwSFlzAAALEwAACxMBAJqcGAAAAAd0SU1FB9wCCBAuLt2rqugAAACMSURBVDjL1ZMxCsJQEERfRPEk9im9mGfxBpaCnaXmBraeQEidKsKz2eIb/peEIOLAwu4Ws7PMbsUAKlOwYCaWmd4aWCV1B4yXpR59R61SitwKDdBHfgPaSTsF8zOmHz5NLykAeCRqvuvCfxGcgP0YF3ZqHy7c1Yt6jfo8dCF3idvkQjcRRVQ/f6bZBC+RBoeZnlCyqwAAAABJRU5ErkJggg==&labelColor=500000)](https://lobste.rs/s/a4jzvv/readme_static_site_riss#c_hiil4z.json)

RISS made it to the first page of [HackerNews](https://news.ycombinator.com/item?id=29304376) and [Lobsters](https://lobste.rs/s/a4jzvv/readme_static_site_riss) and got comments like:

> I never really understood the idea to have a separate README and index page. Glad to see i'm not the only one :)

[southerntofu](https://news.ycombinator.com/item?id=29305519)

> Kudos for making it reusable and not specific to a single static site generator. […]

[hannu](https://lobste.rs/s/a4jzvv/readme_static_site_riss#c_hiil4z)

> […] A small but good idea, I like how simple riss.awk is.

[lifthrasiir](https://news.ycombinator.com/item?id=29305070)

### Run It (Nothing to Install)

To try it with [Hugo][hugo] or [Zola][zola], run the following in your static-site sources:
```sh
wget https://cj.rs/riss.awk
awk -f riss.awk /path/to/my-project/README.md > content/my-project.md
```

If you don’t use Hugo or Zola, no problem! It should also work with any markdown-based static-site generator. Just put the markdown file where it makes sense for that generator.

To automatically update these files in your static-site sources, see [Automate with GitHub Actions](#automate-with-github-actions) below. Since RISS is based on Awk, there is nothing to install, besides copying the script itself!

## Example

### Add a Front Matter

Most static site generators require a “[front matter](https://gohugo.io/getting-started/configuration/#configure-front-matter)” at the beginning of a markdown file to attach some metadata. But you don’t want to add this to your GitHub README! Let’s hide this on GitHub and have it in the script’s output.

In you .md file on GitHub, put:

    <!-- insert
    ---
    title: "README In Static Site (RISS)"
    date: 2021-08-21T10:15:54
    ---
    end_insert -->
    <!-- Powered by https://cj.rs/riss -->
    <!-- remove -->
    
    # README In Static Site (RISS)
    <!-- end_remove -->

The output of the script will be:

    ---
    title: "README In Static Site (RISS)"
    date: 2021-08-21T10:15:54
    ---
    <!-- Powered by https://cj.rs/riss -->

and this piece of yaml will be hidden on GitHub!

### Replace Asciinema Image

You can’t embed the asciinema player on GitHub for security reasons. So the [asciinema documentation](https://asciinema.org/docs/embedding) suggests using an image there and to link it to a webpage with the player. But on your own website, you can embed this player.

In your .md file, put:

    <!-- remove -->
    [![Finding the repositories with “telescope” in their name, with the README in the panel on the right](https://asciinema.org/a/427156.svg)](https://asciinema.org/a/427156)
    <!-- end_remove -->
    <!-- insert
    <asciinema-player src="./telescope.cast" poster="npt:0:04"></asciinema-player>
    end_insert -->

The output will contain only the asciinema player:

    <asciinema-player src="./telescope.cast" poster="npt:0:04"></asciinema-player>

*Note*: you also need to add the JS/CSS files of the asciinema player somewhere in your theme. This [Hugo module][hugo_ascii] makes it easy.

### More

See the [input file (typically on GitHub)](https://github.com/cljoly/readme-in-static-site/blob/main/test.md) and the [output of the script](https://github.com/cljoly/readme-in-static-site/blob/main/test_output.md). You can find another real word [README](https://github.com/cljoly/telescope-repo.nvim/blob/master/README.md) converted to a [webpage](https://cj.rs/telescope-repo-nvim/) (this gives another example of asciinema conversion using a Hugo shortcode).

With some shell scripting, you could also transform all the markdown files in your repo and put them in a subdirectory of your site, so that your project’s documentation, policy, etc… lives on your site or even on a site of its own.

### Your Example!

Have you used this script to transform some markdown (or other) and insert it on your website? [Open an issue][issue] if you would like a link to your use case from this README!

* **telescope-repo.nvim**: [readme](https://github.com/cljoly/telescope-repo.nvim/blob/master/README.md), [website](https://cj.rs/telescope-repo-nvim/); features an Asciinema clip.
* **neovide**: [readme](https://github.com/neovide/neovide/blob/main/README.md), [first iteration of their website](https://neovide.dev/), [PR setting this up](https://github.com/neovide/neovide/pull/1114). They have now moved to mdbook and that’s great! RISS makes the first iteration of your website easy and you are free to move to more complete solutions when your project grows.
* **Hugo APT Repository**: [readme](https://github.com/8hobbies/hugo-apt/blob/master/README.md), [website](https://hugo-apt.8hob.io/), [PR setting it up](https://github.com/8hobbies/hugo-apt/pull/12).

## Transformations Reference

The transformations are driven by HTML comments, so that you can have different results when comments are ignored (e.g. in your GitHub README) and after executing the script on your markdown file.

### Escaping

**It is important that your comment starts at the beginning of the line:** spaces are used for escaping, meaning that if the comment has spaces at the beginning of the line, it is ignored.

So this is escaped
```
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

## Spread the Word

If you find this script useful, please consider inserting the following in your readme:
```html
<!-- Powered by https://cj.rs/riss -->
```
This will help other people find the script. *Thanks for spreading the word!*

If you feel especially charitable, you could put this badge somewhere:

[![](https://img.shields.io/badge/powered%20by-riss-lightgrey)](https://cj.rs/riss)

with for instance this code:
```markdown
[![](https://img.shields.io/badge/powered%20by-riss-lightgrey)](https://cj.rs/riss)
```

## Breaking API Changes

We follow [semver][] and any change that change would cause real world READMEs to be converted differently requires a new major version. In particular, the following is a breaking change:
* adding new keywords (like `remove` or `insert`), as they may be used in the README prior to their introduction in RISS,
* changing a keywords syntax.

## Benchmark

**Processes 17600 lines in 10 ms**

```bash
$ for i in {1..100}; do shuf README.md >>bench.md; done # Create a big md file
$ wc -l README.md
176 README.md
$ wc -l bench.md
17600 bench.md
$ hyperfine 'awk -f riss.awk README.md' 'awk -f riss.awk bench.md'
```

| Command | Mean [ms] | Min [ms] | Max [ms] | Relative |
|:---|---:|---:|---:|---:|
| `awk -f riss.awk README.md` | 2.8 ± 0.2 | 2.4 | 3.7 | 1.00 |
| `awk -f riss.awk bench.md` | 9.7 ± 0.4 | 8.9 | 10.7 | 3.46 ± 0.30 |

## Automate with GitHub Actions

You can automatically update the markdown file in the sources of your site with GitHub Actions. Add this workflow to, for instance, `.github/workflows/readme.yml`:

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

Alternatively, you might configure your repositories to trigger a website rebuild when committing on your README, for instance using [GitHub actions webhooks][gh_webhook].

## Contributions are Welcome!

Feel free to [open an issue][issue] to discuss something or to send a PR.

See also the [Spread the Word][stw] section if you would like to make more folks aware of this script.

![GitHub](https://img.shields.io/github/license/cljoly/readme-in-static-site)

[hugo]: https://gohugo.io/
[zola]: https://www.getzola.org/
[hugo_ascii]: https://cj.rs/gohugo-asciinema/
[issue]: https://github.com/cljoly/readme-in-static-site/issues/new
[fast]: #benchmark
[stw]: #spread-the-word
[script]: https://cj.rs/riss.awk
[gh_webhook]: https://mainawycliffe.dev/blog/github-actions-trigger-via-webhooks/
[semver]: https://semver.org/
