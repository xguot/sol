import Mathlib.Tactic.Linarith

theorem prefix_diff_identity (c1j c2j c1i c2i : ℤ) :
  c1j - c2j = c1i - c2i ↔ c1j - c1i = c2j - c2i := by
  constructor
  all_goals
    intro h
    linarith
