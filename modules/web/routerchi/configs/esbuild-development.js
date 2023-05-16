const esbuild = require("esbuild");
const chokidar = require("chokidar");
const path = require("path");
const shell = require("shelljs");

async function build() {
  const builddir = "tmp";
  const outdir = `${builddir}/web`;
  const entryPoints = ["web/assets/javascripts/application.ts"];
  const watchPaths = [
    "web/templates/**/*.tmpl",
    "web/assets/stylesheets/**/*.css",
    "web/assets/images/**/*.css",
  ];

  shell.rm("-rf", outdir);
  shell.cp("-r", "web/", builddir);

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
  });

  await ctx.watch();
  console.log("Watching...");

  const { host, port } = await ctx.serve({
    servedir: outdir,
    port: 8081,
  });

  console.log(`Listening and serving assets on http://${host}:${port}`);

  const watcher = chokidar.watch(watchPaths, {
    persistent: true,
  });

  watcher.on("add", async (file) => {
    const dirname = `${builddir}/${path.dirname(file)}`;
    const newfile = `${builddir}/${file}`;
    console.log("mkdir -p", dirname);
    shell.mkdir("-p", dirname);
    console.log("cp -r", file, newfile);
    shell.cp(file, newfile);
    await ctx.rebuild();
  });

  watcher.on("change", async (file) => {
    const newfile = `${builddir}/${file}`;
    console.log("cp -r", file, newfile);
    shell.cp(file, newfile);
    await ctx.rebuild();
  });
}

build().catch((e) => {
  console.error(e);
  process.exit(1);
});
