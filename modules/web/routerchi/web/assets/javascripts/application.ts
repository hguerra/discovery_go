import "../stylesheets/application";
import logoUrl from "../images/logo.png";

function main() {
  console.log(">> from js");

  const image = new Image();
  image.src = `${ApplicationConfig.assetsURL}/${logoUrl}`;
  document.body.appendChild(image);
}

document.addEventListener("DOMContentLoaded", main);
