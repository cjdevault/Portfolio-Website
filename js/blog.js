(function () {
  "use strict";

  var cache = new Map();

  function notifyLayoutChange() {
    document.dispatchEvent(new Event("blog-layout-change"));
  }

  function loadPaperContent(body, src) {
    if (cache.has(src)) {
      body.innerHTML = cache.get(src);
      notifyLayoutChange();
      return;
    }

    body.innerHTML = '<p class="blog-post-loading">Loading…</p>';

    fetch(src)
      .then(function (res) {
        if (!res.ok) throw new Error("Failed to load post");
        return res.text();
      })
      .then(function (html) {
        cache.set(src, html);
        body.innerHTML = html;
        notifyLayoutChange();
      })
      .catch(function () {
        body.innerHTML =
          '<p class="blog-post-error">Could not load this paper. Try refreshing the page.</p>';
        notifyLayoutChange();
      });
  }

  document.querySelectorAll(".blog-post").forEach(function (post) {
    var toggle = post.querySelector(":scope > .blog-post-toggle");
    var body = post.querySelector(":scope > .blog-post-body");

    if (!toggle || !body) return;

    toggle.addEventListener("click", function () {
      var isOpen = post.classList.toggle("is-open");
      toggle.setAttribute("aria-expanded", isOpen ? "true" : "false");
      body.hidden = !isOpen;
      notifyLayoutChange();
    });
  });

  document.querySelectorAll(".blog-paper").forEach(function (paper) {
    var toggle = paper.querySelector(".blog-paper-toggle");
    var body = paper.querySelector(".blog-paper-body");
    var src = paper.getAttribute("data-src");

    if (!toggle || !body || !src) return;

    toggle.addEventListener("click", function () {
      var isOpen = paper.classList.toggle("is-open");
      toggle.setAttribute("aria-expanded", isOpen ? "true" : "false");
      body.hidden = !isOpen;
      notifyLayoutChange();

      if (!isOpen) return;

      loadPaperContent(body, src);
    });
  });
})();
