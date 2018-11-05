//package bedtime.moudle.WebFuturelab.Find;
//
//import android.content.Intent;
//import android.os.Bundle;
//import android.support.v4.view.ViewPager;
//import android.view.LayoutInflater;
//import android.view.View;
//import android.view.ViewGroup;
//import android.widget.ImageView;
//import android.widget.LinearLayout;
//import android.widget.TextView;
//
//import com.bedtime.bedtime.R;
//import com.camera.utils.ToastUtils;
//
//import java.util.ArrayList;
//import java.util.List;
//
//import bedtime.common.global.Vardepend;
//import bedtime.common.review.DrawableCenterTextView;
//import bedtime.common.review.tabview.PagerSlidingTabStrip;
//import bedtime.moudle.WebFuturelab.Models.ShowSelectTab;
//import bedtime.moudle.WebFuturelab.WebUtils.UiUtils;
//import bedtime.moudle.bedtimemar.activity.BaseFragment;
//import bedtime.moudle.bedtimemar.activity.FragmentPagerAdapter;
//import bedtime.moudle.login.LoginActivity;
//
///**
// * Created by lijingbo on 2017/12/19.
// */
//
//public class FindFragment extends BaseFragment implements View.OnClickListener {
//    private View layout;
//    private TextView web_tv_release_dynamic_id;
//    private PagerSlidingTabStrip web_psts_find_fragment_id;
//    private ViewPager web_viewpager_find_fragment_id;
//    private ImageView web_iv_release_id;
//    private FragmentPagerAdapter fragmnetpageradapter;
//    private Class [] title_class ;
//    /** 用户栏目列表 */
//    private List<String> titles = new ArrayList<>();//顶部栏名称数组
//    private int numberTable = 3;
//    private LinearLayout msl_vs; //V说那几个标签
//    private DrawableCenterTextView tv_go_tocpic;
//    @Override
//    public View onCreateView(LayoutInflater inflater, ViewGroup container, Bundle savedInstanceState) {
//        layout = inflater.inflate(R.layout.fragment_find_web, container, false);
//        View(layout);
//        return layout;
//    }
//
//    private void View(View layout) {
//
//        web_psts_find_fragment_id = (PagerSlidingTabStrip) layout.findViewById(R.id.web_psts_find_fragment_id);
//        web_psts_find_fragment_id.setOnClickListener(this);
//        web_viewpager_find_fragment_id = (ViewPager) layout.findViewById(R.id.web_viewpager_find_fragment_id);
//        web_viewpager_find_fragment_id.setOnClickListener(this);
//        web_iv_release_id = (ImageView) layout.findViewById(R.id.web_iv_release_id);
//        msl_vs =  layout.findViewById(R.id.msl_vs);
//        tv_go_tocpic = (DrawableCenterTextView) layout.findViewById(R.id.tv_go_tocpic);
//        web_iv_release_id.setOnClickListener(this);
//        tv_go_tocpic.setOnClickListener(this);
//        Bundle[] mBundles = new Bundle[numberTable];
//        title_class =  new Class[numberTable];
//        titles.add("最热");
//        titles.add("近期");
//        titles.add("相关");
//        for (int i = 0; i < numberTable; i++) {
//            title_class[i] = VSTellBottomFragment.class;
//            Bundle mBundle = new Bundle();//1 最热   2 近期    3 相关
//            if (i ==0){
//                mBundle.putString("category_id","1");
//            }else if (i ==1){
//                mBundle.putString("category_id","2");
//            }else if (i ==2){
//                mBundle.putString("category_id","3");
//            }
//            mBundles[i] = mBundle;
//        }
//        fragmnetpageradapter = new FragmentPagerAdapter(getActivity().getSupportFragmentManager(), title_class, titles, mBundles);
//        web_viewpager_find_fragment_id.setAdapter(fragmnetpageradapter);
//        web_psts_find_fragment_id.setViewPager(web_viewpager_find_fragment_id);
//        web_viewpager_find_fragment_id.setOffscreenPageLimit(1);
//        addTableTabCanSkip();
//    }
//
//    @Override
//    public void onClick(View view) {
//        switch (view.getId()) {
//            case R.id.tv_go_tocpic://去查看话题
//                //ActivityUtils.showActivity(getActivity(), .class);
//                break;
//            case R.id.web_iv_release_id:
//                if ("0".equals(Vardepend.getUserid())){
//                    startActivity(new Intent(getActivity(), LoginActivity.class));
//                }else {
//                    Intent intent = new Intent(getActivity(),ReleaseDynamicActivity.class);
//                    intent.putExtra("findType","1");
//                    intent.putExtra("contest_location","");
//                    getActivity().startActivity(intent);
//                }
//                break;
//        }
//    }
//    private void addTableTabCanSkip(){
//        List<ShowSelectTab> list = new ArrayList<>();
//        ShowSelectTab showSelectTab1 = new ShowSelectTab();
//        ShowSelectTab showSelectTab2 = new ShowSelectTab();
//        ShowSelectTab showSelectTab3 = new ShowSelectTab();
//        ShowSelectTab showSelectTab4 = new ShowSelectTab();
//        ShowSelectTab showSelectTab5 = new ShowSelectTab();
//        ShowSelectTab showSelectTab6 = new ShowSelectTab();
//        showSelectTab1.setName("全部");
//        list.add(showSelectTab1);
//        showSelectTab2.setName("智能硬件");
//        list.add(showSelectTab2);
//        showSelectTab3.setName("创新设计");
//        list.add(showSelectTab3);
//        showSelectTab4.setName("人工智能");
//        list.add(showSelectTab4);
//        showSelectTab5.setName("编辑擅长标签");
//        list.add(showSelectTab5);
//        showSelectTab6.setName("编辑兴趣标签");
//        list.add(showSelectTab6);
//        msl_vs.removeAllViews();
//        for (int i = 0 ; i < list.size() ;i++){
//            final ShowSelectTab showSelectTab = list.get(i);
//            UiUtils.addTextViewMyListenerView(showSelectTab, list,msl_vs, new UiUtils.OnMyClickListener() {
//                @Override
//                public void onClick(Object obj) {
//                    ToastUtils.showLongToast(showSelectTab.getName());
//                }
//            });
//        }
//
//    }
//}
