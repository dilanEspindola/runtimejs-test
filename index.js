console.println("Start");

setTimeoutCustom(() => {
  console.println("timeout 1");
}, 4000);

setTimeoutCustom(() => {
  console.println("timeout 2");
}, 100);

function run() {
  for (let i = 0; i < 50; i++) {
    console.println(`item ${i}`);
  }
}

run();
