import { Navigate, Route, Routes } from "react-router-dom";
import Layout from "./components/Layout";
import HomePage from "./pages/HomePage";
import CategoryPage from "./pages/CategoryPage";
import PostDetailPage from "./pages/PostDetailPage";
import NotFoundPage from "./pages/NotFoundPage";
import WritingPage from "./pages/WritingPage";
import PracticePage from "./pages/writing/PracticePage";

export default function App() {
  return (
    <Routes>
      <Route element={<Layout />}>
        <Route index element={<HomePage />} />
        <Route path="tech" element={<CategoryPage category="tech" />} />
        <Route path="gaming" element={<CategoryPage category="gaming" />} />
        <Route path="travelling" element={<CategoryPage category="travelling" />} />
        <Route path="writing" element={<WritingPage />}>
          <Route index element={<Navigate to="practice" replace />} />
          <Route path="practice" element={<PracticePage />} />
        </Route>
        <Route path="posts/:slug" element={<PostDetailPage />} />
        <Route path="*" element={<NotFoundPage />} />
      </Route>
    </Routes>
  );
}
