import { sassPlugin } from 'esbuild-sass-plugin'
import esbuild from 'esbuild'

const args = process.argv.slice(2)
const watch = args.includes('--watch')
const watchPlugin = {
  name: 'watch-plugin',
  setup(build) {
    build.onStart(() => { console.log('Build starting:') })
    build.onEnd((result) => {
      if (result.errors.length > 0) {
        console.log('Build finished with errors')
      } else {
        console.log('Build finished successfully')
      }
    })
  }
}

const context = await esbuild.context( {
    entryPoints: ["frontend/Application.tsx", "frontend/Home.css", "frontend/Navbar.css", "frontend/About.css", "frontend/Flashcard.css", "frontend/Quiz.css"],
    outdir: "public/assets",
    loader: { '.js': 'jsx' },
    bundle: true,
    minify: true,
    plugins: [sassPlugin(), watchPlugin],
})

if (watch) {
  await context.watch()
  console.log('watching...')
} else {
  context.rebuild();
  context.dispose();
}