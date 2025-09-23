import axios from "axios";
import { useState } from "react";

const baseURL = "localhost:8080"

export const api = axios.create({ baseURL: "localhost:8080" })
