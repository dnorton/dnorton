---
layout: default
title: Posts
description: Posts by Daniel Norton
---

<ul class="post-list">
{% for post in site.posts %}
  <li>
    <a href="{{ post.url | relative_url }}">{{ post.title }}</a>
    <p>{{ post.date | date: "%B %-d, %Y" }}</p>
  </li>
{% endfor %}
</ul>
