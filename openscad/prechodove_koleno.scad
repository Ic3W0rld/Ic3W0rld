/* ================================================================
   Prechodové koleno — Rectangular-to-Round Transition Elbow (90°)

   Rectangular opening : rect_w × rect_h  (99 × 236 mm)
   Circular opening    : Ø circ_d         (Ø 160 mm)
   Elbow angle         : 90°

   Orientation:
     - Circular  end at origin  → faces +X  (horizontal)
     - Rectangular end at top   → faces +Z  (vertical)
   ================================================================ */

// ── User Parameters ─────────────────────────────────────────────
rect_w   = 99;    // rectangular duct width  [mm]
rect_h   = 236;   // rectangular duct height [mm]
circ_d   = 160;   // circular duct diameter  [mm]

wall_t   = 1.5;   // wall thickness          [mm]
inner_r  = 20;    // inner bend radius       [mm]
stub_len = 25;    // straight stub at each opening [mm]
cr_min   = 2;     // minimum corner radius at rectangular end [mm]

// ── Derived ─────────────────────────────────────────────────────
// Centerline arc radius — bend is in X-Z plane, so the in-plane
// half-dimensions at each end are: rect_w/2 (rect end) and circ_d/2 (circ end).
// Use the larger to ensure clearance at both openings.
R  = inner_r + max(circ_d / 2, rect_w / 2);   // = 100 mm with defaults

NS = 24;    // number of segments in the 90° bend
$fn = 64;

// ── Cross-section profile ───────────────────────────────────────
// t = 0  →  circular  (Ø circ_d)
// t = 1  →  rectangular (rect_w × rect_h, corners ~cr_min)
//
// With extra > 0 the profile expands uniformly by `extra` on all sides
// (used to add wall thickness to the outer shell).
module cs(t, extra = 0) {
    w  = circ_d * (1-t) + rect_w * t;
    h  = circ_d * (1-t) + rect_h * t;
    cr = cr_min + (circ_d/2 - cr_min) * (1 - t);   // smoothly decreasing corner radius

    offset(r = cr + extra, $fn = 64)
        square([max(0.01, w - 2*cr), max(0.01, h - 2*cr)], center = true);
}

// ── Thin positioning disk along the quarter-circle arc ──────────
// Arc: center at world (0, 0, R)
//   theta = 0°  →  position (0,   0, 0),   face +X  (circular end)
//   theta = 90° →  position (R,   0, R),   face +Z  (rectangular end)
module disk(theta, extra = 0) {
    t = theta / 90;
    x = R * sin(theta);
    z = R * (1 - cos(theta));

    translate([x, 0, z])
    rotate([0, 90 - theta, 0])          // orient face tangent to arc
    linear_extrude(height = 0.01, center = true)
        cs(t, extra);
}

// ── Hollow 90° bend ─────────────────────────────────────────────
module bend() {
    difference() {
        // outer wall
        union() {
            for (i = [0 : NS-1]) {
                hull() {
                    disk(i       / NS * 90, wall_t);
                    disk((i + 1) / NS * 90, wall_t);
                }
            }
        }
        // inner void
        union() {
            for (i = [0 : NS-1]) {
                hull() {
                    disk(i       / NS * 90);
                    disk((i + 1) / NS * 90);
                }
            }
        }
    }
}

// ── Straight hollow stub ─────────────────────────────────────────
// Extrudes cross-section cs(t) by `length` along local +Z.
module stub(t, length) {
    difference() {
        linear_extrude(length)
            cs(t, wall_t);
        translate([0, 0, -0.1])
        linear_extrude(length + 0.2)
            cs(t);
    }
}

// ── Final assembly ───────────────────────────────────────────────
module transition_elbow() {
    // 90° elbow body
    bend();

    // Circular stub  — extends in +X from origin
    rotate([0, -90, 0])
        stub(0, stub_len);

    // Rectangular stub — extends in +Z from top of arc
    translate([R, 0, R])
        stub(1, stub_len);
}

transition_elbow();
