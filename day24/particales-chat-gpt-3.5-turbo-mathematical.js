// Function to solve a system of linear equations using Gaussian elimination
function solveLinearEquations(A, B) {
    const n = A.length;
    const augmentedMatrix = A.map((row, index) => [...row, B[index]]);
  
    for (let i = 0; i < n; i++) {
      const divisor = augmentedMatrix[i][i];
      for (let j = i; j < n + 1; j++) {
        augmentedMatrix[i][j] /= divisor;
      }
  
      for (let k = 0; k < n; k++) {
        if (k !== i) {
          const factor = augmentedMatrix[k][i];
          for (let j = i; j < n + 1; j++) {
            augmentedMatrix[k][j] -= factor * augmentedMatrix[i][j];
          }
        }
      }
    }
  
    return augmentedMatrix.map(row => row[n]);
  }
  
  // Function to check if two hailstones will intersect within a specified test area
  function willHailstonesIntersect(hailstoneA, hailstoneB, testArea) {
    const A = [
      [hailstoneA.velocity.x, -hailstoneB.velocity.x],
      [hailstoneA.velocity.y, -hailstoneB.velocity.y]
    ];
  
    const B = [
      hailstoneB.position.x - hailstoneA.position.x,
      hailstoneB.position.y - hailstoneA.position.y
    ];
  
    const [t1, t2] = solveLinearEquations(A, B);
  
    if (t1 >= 0 && t2 >= 0) {
      // Calculate intersection points
      const intersectionX = hailstoneA.position.x + hailstoneA.velocity.x * t1;
      const intersectionY = hailstoneA.position.y + hailstoneA.velocity.y * t1;
  
      // Check if the intersection points are within the test area
      return (
        intersectionX >= testArea.minX &&
        intersectionX <= testArea.maxX &&
        intersectionY >= testArea.minY &&
        intersectionY <= testArea.maxY
      );
    }
  
    return false;
  }
  
  // Function to find the number of intersections within a test area
  function findIntersectionsCount(hailstones, testArea) {
    let count = 0;
  
    for (let i = 0; i < hailstones.length - 1; i++) {
      for (let j = i + 1; j < hailstones.length; j++) {
        const intersect = willHailstonesIntersect(hailstones[i], hailstones[j], testArea);
        if (intersect) {
          count++;
        }
      }
    }
  
    return count;
  }
  
  // Example usage with the given input
  const hailstones = [
    { position: { x: 19, y: 13 }, velocity: { x: -2, y: 1 } },
    { position: { x: 18, y: 19 }, velocity: { x: -1, y: -1 } },
    { position: { x: 20, y: 25 }, velocity: { x: -2, y: -2 } },
    { position: { x: 12, y: 31 }, velocity: { x: -1, y: -2 } },
    { position: { x: 20, y: 19 }, velocity: { x: 1, y: -5 } }
  ];
  
  const testArea = {
    minX: 7,
    maxX: 27,
    minY: 7,
    maxY: 27
  };
  
  const intersectionsCount = findIntersectionsCount(hailstones, testArea);
  console.log(intersectionsCount); // Output the number of intersections
  