(function () {
  "use strict";

  const PANEL_ORDER = ["about", "skills", "projects", "resume", "contact"];

  const navItems = document.querySelectorAll(".nav-item");
  const panels = document.querySelectorAll(".panel");
  const bgSlides = document.querySelectorAll(".bg-slide");
  const prevBtn = document.querySelector(".nav-arrow--prev");
  const nextBtn = document.querySelector(".nav-arrow--next");

  let currentIndex = 0;

  function panelIndex(id) {
    return PANEL_ORDER.indexOf(id);
  }

  function showPanel(panelId) {
    const index = panelIndex(panelId);
    if (index < 0) return;

    currentIndex = index;

    navItems.forEach(function (btn) {
      const isActive = btn.getAttribute("data-panel") === panelId;
      btn.classList.toggle("active", isActive);
      btn.setAttribute("aria-current", isActive ? "page" : "false");
    });

    panels.forEach(function (panel) {
      const isActive = panel.getAttribute("data-panel") === panelId;
      panel.classList.toggle("active", isActive);
      panel.hidden = !isActive;
    });

    bgSlides.forEach(function (slide) {
      slide.classList.toggle("active", slide.getAttribute("data-panel") === panelId);
    });

    document.body.setAttribute("data-panel", panelId);

    if (prevBtn) prevBtn.disabled = index === 0;
    if (nextBtn) nextBtn.disabled = index === PANEL_ORDER.length - 1;
  }

  function showByIndex(index) {
    if (index < 0 || index >= PANEL_ORDER.length) return;
    showPanel(PANEL_ORDER[index]);
  }

  navItems.forEach(function (btn) {
    btn.addEventListener("click", function () {
      showPanel(btn.getAttribute("data-panel"));
    });
  });

  if (prevBtn) {
    prevBtn.addEventListener("click", function () {
      showByIndex(currentIndex - 1);
    });
  }

  if (nextBtn) {
    nextBtn.addEventListener("click", function () {
      showByIndex(currentIndex + 1);
    });
  }

  document.addEventListener("keydown", function (e) {
    if (e.key === "ArrowLeft") {
      showByIndex(currentIndex - 1);
    } else if (e.key === "ArrowRight") {
      showByIndex(currentIndex + 1);
    }
  });

  showPanel("about");
})();
