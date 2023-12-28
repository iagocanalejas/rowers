/**
 * @param {boolean} visible - Content visibility
 */
function toggleVisibility(visible) {
  let component = document.getElementById("collapsibleContent");
  let button = document.getElementById("collapsibleButton");
  if (component) {
    if (visible) {
      component.classList.remove("hidden");
      button.classList.add("hidden");
    } else {
      component.classList.add("hidden");
      button.classList.remove("hidden");
    }
  }
}
