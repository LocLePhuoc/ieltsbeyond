---
title: "Why I Switched to Neovim"
summary: "After years of VS Code, I made the jump to Neovim. Here's what convinced me and what I learned along the way."
date: "2026-01-15"
tags: ["neovim", "editor", "productivity"]
cover_image: ""
---

I'll be honest — the first time I tried Vim, I couldn't figure out how to quit. The classic meme was my lived experience. But something kept pulling me back.

## The Breaking Point

VS Code had been my daily driver for four years. It's a phenomenal editor, and I have nothing bad to say about it. But as my projects grew larger and my workflow became more keyboard-driven, I kept bumping into friction. The mouse felt like an interruption. Extensions loaded slowly. Remote development over SSH was workable but clunky.

Then I watched a colleague navigate a large codebase in Neovim. No mouse. No pauses. Just thought translated directly into action. It looked like magic, but it was really just muscle memory and a well-configured tool.

## The Learning Curve

Let's not sugarcoat it — the learning curve is real. The first two weeks were painful. My productivity dropped by at least 50%. I kept a cheat sheet taped to my monitor and resisted the urge to open VS Code "just for this one thing."

The breakthrough came when I stopped thinking of Vim motions as commands and started thinking of them as a **language**. `d2w` isn't "press d, then 2, then w" — it's "delete two words." Once that mental model clicked, everything accelerated.

## My Setup Today

I use **lazy.nvim** for plugin management, **Telescope** for fuzzy finding, **nvim-lspconfig** for language servers, and **Treesitter** for syntax highlighting. The configuration lives in Lua, version-controlled in a dotfiles repo that follows me across machines.

The entire setup starts in under 100 milliseconds. Every keystroke does exactly what I expect. And when it doesn't, I can fix it — because it's all just configuration.

Is it for everyone? Absolutely not. But for me, it turned editing code from something I did into something I *enjoyed*.
