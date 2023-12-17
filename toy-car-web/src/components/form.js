// src/components/RecipeForm.js
import React, { useState } from 'react';

function PlacementForm({ placeCar, carPosition }) {

  const [position, setPosition] = useState({x: 0, y: 0, direction: "NORTH" });

  const handleInputChange = (event) => {
    let { name, value } = event.target;
    if (name !== "direction") {
      value = parseInt(value, 10)
    }
    setPosition({ ...position, [name]: value });
  };

  const handleSubmit = (event) => {
    event.preventDefault();
    setPosition({ ...position })
    placeCar({ ...position });
  };

  return (
    <div>
      <form onSubmit={handleSubmit}>
        <input
          type="text"
          name="x"
          placeholder="Posição X"
          value={position.x}
          onChange={handleInputChange}
        />
        <input
          type="text"
          name="y"
          placeholder="Position Y"
          value={position.y}
          onChange={handleInputChange}
        />
        <select
          name="direction"
          value={position.direction}
          onChange={handleInputChange}
        >
          <option value="SOUTH">South</option>
          <option value="NORTH">North</option>
          <option value="EAST">East</option>
          <option value="WEST">West</option>
        </select>
        <button type="submit">Posicionar Carro</button>
      </form>
    </div>
  );
}

export default PlacementForm;
