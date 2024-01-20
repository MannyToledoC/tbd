import styles from "./page.module.css";
import Test from "@/Components/Test/Test";

export default function Home() {
  return (
    <main>
      <div className={styles.description}>
        <Test />
        Hello
      </div>
    </main>
  );
}
