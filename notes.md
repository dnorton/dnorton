---
layout: default
title: Notes
description: Notes by Daniel Norton
permalink: /notes/
---

<ul class="post-list">
{% assign writing = site.writing | sort: 'date' | reverse %}
{% for post in writing %}
  <li>
    <a href="{{ post.url | relative_url }}">{{ post.title }}</a>
    <p>{{ post.date | date: "%B %-d, %Y" }}</p>
  </li>
{% endfor %}
</ul>
