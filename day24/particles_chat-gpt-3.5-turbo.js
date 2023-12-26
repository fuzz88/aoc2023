function parseInput(input) {
    return input
      .trim()
      .split('\n')
      .map((line) => {
        const [px, py, pz, vx, vy, vz] = line
          .match(/(-?\d+)/g)
          .map((num) => parseInt(num, 10));
        return { position: { x: px, y: py, z: pz }, velocity: { x: vx, y: vy, z: vz } };
      });
  }
  
  function simulateHailstones(hailstones, steps) {
    return hailstones.map(({ position, velocity }) => ({
      position: {
        x: position.x + velocity.x * steps,
        y: position.y + velocity.y * steps,
        z: position.z + velocity.z * steps,
      },
      velocity,
    }));
  }
  
  function arePathsIntersecting(a, b) {
    return a.position.x === b.position.x && a.position.y === b.position.y;
  }
  
  function findIntersections(hailstones) {
    const intersections = [];
  
    for (let i = 0; i < hailstones.length - 1; i++) {
      for (let j = i + 1; j < hailstones.length; j++) {
        const hailstoneA = hailstones[i];
        const hailstoneB = hailstones[j];
  
        if (arePathsIntersecting(hailstoneA, hailstoneB)) {
          intersections.push({ hailstoneA, hailstoneB });
        }
      }
    }
  
    return intersections;
  }
  
  function printIntersections(intersections) {
    intersections.forEach(({ hailstoneA, hailstoneB }) => {
      console.log(`Hailstone A: ${hailstoneA.position.x}, ${hailstoneA.position.y}`);
      console.log(`Hailstone B: ${hailstoneB.position.x}, ${hailstoneB.position.y}`);
      console.log('Hailstones\' paths crossed in the past for both hailstones.');
      console.log();
    });
  }
  
  const input = `19, 13, 30 @ -2, 1, -2
  18, 19, 22 @ -1, -1, -2
  20, 25, 34 @ -2, -2, -4
  12, 31, 28 @ -1, -2, -1
  20, 19, 15 @ 1, -5, -3`;
  
  const hailstones = parseInput(input);
  
  const testAreaMinX = 7;
  const testAreaMaxX = 27;
  const testAreaMinY = 7;
  const testAreaMaxY = 27;
  
  const steps = 1000; // Adjust the number of steps based on your input data
  
  for (let step = 1; step <= steps; step++) {
    const simulatedHailstones = simulateHailstones(hailstones, step);
    const intersections = findIntersections(simulatedHailstones);
  
    intersections.forEach(({ hailstoneA, hailstoneB }) => {
      console.log(`Hailstone A: ${hailstoneA.position.x}, ${hailstoneA.position.y}`);
      console.log(`Hailstone B: ${hailstoneB.position.x}, ${hailstoneB.position.y}`);
      console.log(`Hailstones' paths will cross inside the test area (at x=${(hailstoneA.position.x + hailstoneB.position.x) / 2}, y=${(hailstoneA.position.y + hailstoneB.position.y) / 2}).`);
      console.log();
    });
  }
  