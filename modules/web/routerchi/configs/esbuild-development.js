const esbuild = require("esbuild");
const copyStaticFiles = require("esbuild-copy-static-files");

async function build() {
  const builddir = "tmp";
  const outdir = `${builddir}/web`;
  const entryPoints = ["web/assets/javascripts/application.ts"];

  const ctx = await esbuild.context({
    entryPoints,
    outdir,
    outbase: "web",
    platform: "browser",
    format: "iife",
    bundle: true,
    metafile: false,
    minify: false,
    sourcemap: true,
    treeShaking: false,
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
    ],
  });

  await ctx.watch();
  console.log("Watching...");

  const { host, port } = await ctx.serve({
    servedir: outdir,
    port: 8081,
  });

  console.log(`Listening and serving assets on http://${host}:${port}`);
}

build().catch((e) => {
  console.error(e);
  process.exit(1);
});
