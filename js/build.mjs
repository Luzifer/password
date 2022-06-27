import { sassPlugin } from 'esbuild-sass-plugin'
import esbuild from 'esbuild'

esbuild.build({
  bundle: true,
  define: {
    'process.env.NODE_ENV': JSON.stringify(process.env.NODE_ENV || 'dev'),
  },
  entryPoints: ['src/index.js'],
  loader: {
    '.woff': 'file',
    '.woff2': 'file',
  },
  minify: true,
  outfile: '../cmd/password/frontend/assets/bundle.js',
  plugins: [sassPlugin()],
  target: [
    'chrome87',
    'edge87',
    'es2020',
    'firefox84',
    'safari14',
  ],
})
