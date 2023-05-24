const esbuild = require("esbuild");
const copyStaticFiles = require("esbuild-copy-static-files");
const manifestPlugin = require("esbuild-plugin-manifest");

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
    plugins: [
      copyStaticFiles({
        src: "web/assets/images",
        dest: `${outdir}/assets/images`,
        dereference: true,
        errorOnExist: false,
        preserveTimestamps: true,
        recursive: true,
      }),
      copyStaticFiles({
        src: "web/public",
        dest: `${outdir}/public`,
        dereference: true,
        errorOnExist: false,
        preserveTimestamps: true,
        recursive: true,
      }),
      copyStaticFiles({
        src: "web/templates",
        dest: `${outdir}/templates`,
        dereference: true,
        errorOnExist: false,
        preserveTimestamps: true,
        recursive: true,
      }),
      manifestPlugin(),
    ],
  });

  if (result.errors.length) {
    throw new Error(`Fail to build: ${result.errors}`);
  }

  console.log(
    await esbuild.analyzeMetafile(result.metafile, {
      verbose: true,
    })
  );
}

build().catch((e) => {
  console.error(e);
  process.exit(1);
});
