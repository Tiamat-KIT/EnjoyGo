mkdir dist
cp "$(tinygo env TINYGOROOT)/targets/wasm_exec.js" ./dist/wasm_exec.js
tinygo build -o ./dist/output.wasm -target wasm ./main.go