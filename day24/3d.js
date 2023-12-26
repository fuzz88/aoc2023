function findCollisionPositionAndVelocity(hailstones) {
    // Function to check if the rock collides with all hailstones at a given time
    function checkCollision(time, position, velocity) {
      for (const hailstone of hailstones) {
        const projectedPosition = {
          x: hailstone.position.x + hailstone.velocity.x * time,
          y: hailstone.position.y + hailstone.velocity.y * time,
          z: hailstone.position.z + hailstone.velocity.z * time
        };
  
        // Check if the rock collides with the hailstone at the projected position
        if (
          Math.abs(projectedPosition.x - position.x) <= 1 &&
          Math.abs(projectedPosition.y - position.y) <= 1 &&
          Math.abs(projectedPosition.z - position.z) <= 1
        ) {
          continue;
        } else {
          return false;
        }
      }
      return true;
    }
  
    // Iterate through possible positions and velocities to find a solution
    for (let x = 0; x <= 1000; x++) {
      for (let y = 0; y <= 1000; y++) {
        for (let z = 0; z <= 1000; z++) {
          for (let vx = -10; vx <= 10; vx++) {
            for (let vy = -10; vy <= 10; vy++) {
              for (let vz = -10; vz <= 10; vz++) {
                const position = { x, y, z };
                const velocity = { x: vx, y: vy, z: vz };
  
                let collisionTime = null;
                for (let time = 0; time <= 100; time++) {
                  if (checkCollision(time, position, velocity)) {
                    collisionTime = time;
                    break;
                  }
                }
  
                // If a collision time is found, return the result
                if (collisionTime !== null) {
                  return {
                    position,
                    velocity
                  };
                }
              }
            }
          }
        }
      }
    }
  
    // If no solution is found, return null
    return null;
  }
  
  // Example usage with the given hailstones
  const hailstones = [
    { position: { x: 19, y: 13, z: 30 }, velocity: { x: -2, y: 1, z: -2 } },
    { position: { x: 18, y: 19, z: 22 }, velocity: { x: -1, y: -1, z: -2 } },
    { position: { x: 20, y: 25, z: 34 }, velocity: { x: -2, y: -2, z: -4 } },
    { position: { x: 12, y: 31, z: 28 }, velocity: { x: -1, y: -2, z: -1 } },
    { position: { x: 20, y: 19, z: 15 }, velocity: { x: 1, y: -5, z: -3 } }
  ];
  
  const result = findCollisionPositionAndVelocity(hailstones);
  console.log(result.position.x + result.position.y + result.position.z); // Output the sum of X, Y, and Z coordinates
  