import { getCredentials } from "@/utils/auth";
import React from "react";
import { Redirect, Route } from "wouter";

type Props = React.ComponentProps<typeof Route>;

const ProtectedRoute = (props: Props) => {
  const token = getCredentials().acces_token;

  if(!token) {
    return <Redirect to="/login" />;
  }

  return <Route {...props} />;
};

export default ProtectedRoute;
