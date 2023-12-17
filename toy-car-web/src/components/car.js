import React from "react";
import './car.css'

function Car({ direction }) {
  return <div className={`car-${direction}`}>🚗</div>;
}

export default Car
