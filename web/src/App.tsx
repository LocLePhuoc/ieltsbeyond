import { Navigate, Route, Routes } from "react-router-dom";
import Layout from "./components/Layout";
import NotFoundPage from "./pages/NotFoundPage";
import WritingPage from "./pages/WritingPage";
import OverviewPage from "./pages/writing/OverviewPage";
import PracticePage from "./pages/writing/PracticePage";

export default function App() {
  return (
    <Routes>
      <Route element={<Layout />}>
        <Route index element={<Navigate to="/writing" replace />} />
        <Route path="writing" element={<WritingPage />}>
          <Route index element={<OverviewPage />} />
          <Route path="practice" element={<PracticePage />} />
        </Route>
        <Route path="*" element={<NotFoundPage />} />
      </Route>
    </Routes>
  );
}
