import { Route, Routes } from "react-router-dom";
import Layout from "./components/Layout";
import HomePage from "./pages/HomePage";
import CategoryPage from "./pages/CategoryPage";
import PostDetailPage from "./pages/PostDetailPage";
import NotFoundPage from "./pages/NotFoundPage";
import AdminPostListPage from "./admin/AdminPostListPage";
import AdminPostEditorPage from "./admin/AdminPostEditorPage";

export default function App() {
  return (
    <Routes>
      <Route path="admin/posts" element={<AdminPostListPage />} />
      <Route path="admin/posts/new" element={<AdminPostEditorPage />} />
      <Route path="admin/posts/:id/edit" element={<AdminPostEditorPage />} />
      <Route element={<Layout />}>
        <Route index element={<HomePage />} />
        <Route path="tech" element={<CategoryPage category="tech" />} />
        <Route path="gaming" element={<CategoryPage category="gaming" />} />
        <Route path="travelling" element={<CategoryPage category="travelling" />} />
        <Route path="posts/:slug" element={<PostDetailPage />} />
        <Route path="*" element={<NotFoundPage />} />
      </Route>
    </Routes>
  );
}
