(function () {
  "use strict";

  if (!document.body.classList.contains("subpage")) return;

  var NAV_ZONE = 72;

  function updateTopNavBacking() {
    var anchor = document.querySelector(".blog-posts") || document.querySelector(".blog-main");
    if (!anchor) return;

    var top = anchor.getBoundingClientRect().top;
    document.body.classList.toggle("top-nav-backed", top < NAV_ZONE);
  }

  window.addEventListener("scroll", updateTopNavBacking, { passive: true });
  window.addEventListener("resize", updateTopNavBacking);
  document.addEventListener("blog-layout-change", updateTopNavBacking);

  updateTopNavBacking();
})();
