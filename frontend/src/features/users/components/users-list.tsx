import { useQuery } from "@tanstack/react-query";
import React from "react";
import { getAllUsers } from "../services/users.service";

const UsersList = () => {
  const { data } = useQuery({
    queryKey: ["users"],
    queryFn: getAllUsers,
  });

  return <ul>
    {data?.map((user) => (
      <li key={user.id}>{user.username}</li>
    ))}
  </ul>;
};

export default UsersList;
