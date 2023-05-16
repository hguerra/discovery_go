const esbuild = require("esbuild");
const fs = require("fs");

async function build() {
  const outdir = "build/web";
  const entryPoints = ["web/assets/javascripts/application.ts"];

  const result = await esbuild.build({
    entryPoints,
    outdir,
    outbase: "web",
    entryNames: "[dir]/[name]-[hash]",
    platform: "browser",
    format: "iife",
    bundle: true,
    metafile: true,
    minify: true,
    sourcemap: false,
    treeShaking: true,
    loader: {
      ".ts": "ts",
      ".tsx": "tsx",
      ".css": "css",
      ".png": "file",
      ".svg": "file",
      ".html": "text",
    },
  });

  if (result.errors.length) {
    throw new Error(`Fail to build: ${result.errors}`);
  }

  console.log(
    await esbuild.analyzeMetafile(result.metafile, {
      verbose: true,
    })
  );

  fs.writeFileSync(`${outdir}/meta.json`, JSON.stringify(result.metafile));
}

build().catch((e) => {
  console.error(e);
  process.exit(1);
});
