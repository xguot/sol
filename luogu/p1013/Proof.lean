import Mathlib

open Finset

theorem carry_count_eq_value (B x : ℕ) (hx : x < B) :
  (filter (fun y => B ≤ x + y) (range B)).card = x := by
  have h : filter (fun y => B ≤ x + y) (range B) = Ico (B - x) B := by
    ext y
    simp only [mem_filter, mem_range, mem_Ico]
    omega
  rw [h, Nat.card_Ico]
  omega
