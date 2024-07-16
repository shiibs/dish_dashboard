import { useEffect, useState } from "react";
import axios from "axios";
import {
  Card,
  CardHeader,
  CardBody,
  CardFooter,
  Typography,
  Button,
} from "@material-tailwind/react";

export default function Dishes() {
  const [dishes, setDishes] = useState([]);

  useEffect(() => {
    fetchDishes();
    connectWebSocket();
  }, []);

  // Connects to the WebSocket server and handles incoming messages
  const connectWebSocket = () => {
    const ws = new WebSocket("ws://localhost:8005/ws");

    ws.onopen = () => {
      console.log("Connected to WebSocket");
    };

    ws.onmessage = (event) => {
      const updatedDish = JSON.parse(event.data);
      setDishes((prevDishes) =>
        prevDishes.map((dish) =>
          dish.dishId === updatedDish.dishId ? updatedDish : dish
        )
      );
    };

    ws.onclose = () => {
      console.log("Disconnected from WebSocket");
    };

    ws.onerror = (error) => {
      console.error("WebSocket error:", error);
    };
  };

  // Fetches the list of dishes from the server
  const fetchDishes = async () => {
    try {
      const response = await axios.get("http://localhost:8005/dishes");
      setDishes(response.data);
    } catch (err) {
      console.log(err);
    }
  };

  // Toggles the publish status of a dish
  const togglePublish = async (id) => {
    try {
      const response = await axios.put(
        "http://localhost:8005/dishes/toggle/" + id
      );
      setDishes((prevDishes) =>
        prevDishes.map((dish) => (dish.dishId === id ? response.data : dish))
      );
    } catch (err) {
      console.log(err);
    }
  };

  return (
    <div className="container mx-auto p-4">
      <ul className="grid grid-cols-3 gap-4">
        {dishes.map((dish) => (
          <li key={dish.dishId}>
            <Card className="mt-6 w-96 mb-8">
              <CardHeader color="blue-gray" className="relative h-56">
                <img src={dish.imageUrl} alt={dish.dishName} />
              </CardHeader>
              <CardBody>
                <Typography variant="h5" color="blue-gray" className="mb-2">
                  {dish.dishName}
                </Typography>
              </CardBody>
              <CardFooter className="pt-0">
                <Button
                  onClick={() => togglePublish(dish.dishId)}
                  className={`${
                    dish.isPublished
                      ? "bg-green-500 hover:bg-green-700"
                      : "bg-blue-500 hover:bg-blue-700"
                  } text-white font-bold py-2 px-4 rounded`}
                >
                  {dish.isPublished ? "Unpublish" : "Publish"}
                </Button>
              </CardFooter>
            </Card>
          </li>
        ))}
      </ul>
    </div>
  );
}
