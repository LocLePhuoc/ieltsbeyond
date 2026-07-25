import { EditorContent, useEditor } from "@tiptap/react";
import StarterKit from "@tiptap/starter-kit";
import Placeholder from "@tiptap/extension-placeholder";
import Underline from "@tiptap/extension-underline";
import Link from "@tiptap/extension-link";
import TextAlign from "@tiptap/extension-text-align";
import Highlight from "@tiptap/extension-highlight";
import type { JSONContent } from "@tiptap/core";

const emptyDoc: JSONContent = {
  type: "doc",
  content: [{ type: "paragraph" }],
};

interface TiptapEditorProps {
  content: JSONContent | null;
  onChange: (json: JSONContent, html: string) => void;
}

function ToolbarButton({ label, active, disabled, onClick, title }: { label: string; active?: boolean; disabled?: boolean; onClick?: () => void; title?: string }) {
  return (
    <button
      type="button"
      title={title ?? label}
      disabled={disabled}
      onClick={onClick}
      className={active ? "simple-editor-button simple-editor-button-active" : "simple-editor-button"}
    >
      {label}
    </button>
  );
}

function Divider() {
  return <span className="simple-editor-divider" />;
}

export default function TiptapEditor({ content, onChange }: TiptapEditorProps) {
  const editor = useEditor({
    extensions: [
      StarterKit.configure({
        heading: { levels: [1, 2, 3] },
      }),
      Underline,
      Link.configure({
        openOnClick: false,
        autolink: true,
      }),
      TextAlign.configure({
        types: ["heading", "paragraph"],
      }),
      Highlight.configure({ multicolor: false }),
      Placeholder.configure({ placeholder: "Start writing..." }),
    ],
    content: content ?? emptyDoc,
    immediatelyRender: false,
    editorProps: {
      attributes: {
        class: "simple-editor-content focus:outline-none",
      },
    },
    onUpdate: ({ editor }) => onChange(editor.getJSON(), editor.getHTML()),
  });

  if (!editor) return null;

  const setLink = () => {
    const previousUrl = editor.getAttributes("link").href as string | undefined;
    const url = window.prompt("URL", previousUrl ?? "https://");
    if (url === null) return;
    if (url === "") {
      editor.chain().focus().extendMarkRange("link").unsetLink().run();
      return;
    }
    editor.chain().focus().extendMarkRange("link").setLink({ href: url }).run();
  };

  return (
    <div className="simple-editor-shell">
      <div className="simple-editor-toolbar">
        <ToolbarButton label="↶" title="Undo" onClick={() => editor.chain().focus().undo().run()} />
        <ToolbarButton label="↷" title="Redo" onClick={() => editor.chain().focus().redo().run()} />
        <Divider />

        <select
          aria-label="Text style"
          className="simple-editor-select"
          value={editor.isActive("heading", { level: 1 }) ? "h1" : editor.isActive("heading", { level: 2 }) ? "h2" : editor.isActive("heading", { level: 3 }) ? "h3" : "p"}
          onChange={(event) => {
            const value = event.target.value;
            if (value === "p") editor.chain().focus().setParagraph().run();
            if (value === "h1") editor.chain().focus().toggleHeading({ level: 1 }).run();
            if (value === "h2") editor.chain().focus().toggleHeading({ level: 2 }).run();
            if (value === "h3") editor.chain().focus().toggleHeading({ level: 3 }).run();
          }}
        >
          <option value="p">Text</option>
          <option value="h1">Heading 1</option>
          <option value="h2">Heading 2</option>
          <option value="h3">Heading 3</option>
        </select>

        <ToolbarButton label="•≡" title="Bullet list" active={editor.isActive("bulletList")} onClick={() => editor.chain().focus().toggleBulletList().run()} />
        <ToolbarButton label="1≡" title="Numbered list" active={editor.isActive("orderedList")} onClick={() => editor.chain().focus().toggleOrderedList().run()} />
        <ToolbarButton label="❝" title="Quote" active={editor.isActive("blockquote")} onClick={() => editor.chain().focus().toggleBlockquote().run()} />
        <Divider />

        <ToolbarButton label="B" title="Bold" active={editor.isActive("bold")} onClick={() => editor.chain().focus().toggleBold().run()} />
        <ToolbarButton label="I" title="Italic" active={editor.isActive("italic")} onClick={() => editor.chain().focus().toggleItalic().run()} />
        <ToolbarButton label="S" title="Strike" active={editor.isActive("strike")} onClick={() => editor.chain().focus().toggleStrike().run()} />
        <ToolbarButton label="</>" title="Inline code" active={editor.isActive("code")} onClick={() => editor.chain().focus().toggleCode().run()} />
        <ToolbarButton label="U" title="Underline" active={editor.isActive("underline")} onClick={() => editor.chain().focus().toggleUnderline().run()} />
        <ToolbarButton label="✎" title="Highlight" active={editor.isActive("highlight")} onClick={() => editor.chain().focus().toggleHighlight().run()} />
        <ToolbarButton label="🔗" title="Link" active={editor.isActive("link")} onClick={setLink} />
        <Divider />

        <ToolbarButton label="☰" title="Align left" active={editor.isActive({ textAlign: "left" })} onClick={() => editor.chain().focus().setTextAlign("left").run()} />
        <ToolbarButton label="≡" title="Align center" active={editor.isActive({ textAlign: "center" })} onClick={() => editor.chain().focus().setTextAlign("center").run()} />
        <ToolbarButton label="☷" title="Align right" active={editor.isActive({ textAlign: "right" })} onClick={() => editor.chain().focus().setTextAlign("right").run()} />
        <ToolbarButton label="☵" title="Justify" active={editor.isActive({ textAlign: "justify" })} onClick={() => editor.chain().focus().setTextAlign("justify").run()} />
        <Divider />

        <ToolbarButton label="{}" title="Code block" active={editor.isActive("codeBlock")} onClick={() => editor.chain().focus().toggleCodeBlock().run()} />
      </div>

      <EditorContent editor={editor} />
    </div>
  );
}
